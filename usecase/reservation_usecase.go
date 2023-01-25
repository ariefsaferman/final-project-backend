package usecase

import (
	"time"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/constant"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/repository"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
)

type ReservationUseCase interface {
	CreateReservation(req *dto.ReservationRequest, userId uint) (*entity.Reservation, error)
}

type reservationUseCaseImpl struct {
	reservationRepo repository.ReservationRepositry
	walletRepo      repository.WalletRepository
	houseRepo       repository.HouseRepository
}

type ReservationUConfig struct {
	ReservationRepo repository.ReservationRepositry
	WalletRepo      repository.WalletRepository
	HouseRepo       repository.HouseRepository
}

func NewReservationUseCase(config *ReservationUConfig) ReservationUseCase {
	return &reservationUseCaseImpl{
		reservationRepo: config.ReservationRepo,
		walletRepo:      config.WalletRepo,
		houseRepo:       config.HouseRepo,
	}
}

func (r *reservationUseCaseImpl) CreateReservation(req *dto.ReservationRequest, userId uint) (*entity.Reservation, error) {
	reservation := req.ReqToReservation(userId)
	house, err := r.houseRepo.GetHouseById(reservation.HouseID)
	pickup := &entity.PickUp{}
	hostId := house.UserID
	if err != nil {
		return nil, err
	}

	if req.CityID != nil {
		pickup.UserID = reservation.UserID
		pickup.PickupStatusID = constant.PENDING
	}

	if err = r.reservationRepo.CheckVacancy(reservation.HouseID, reservation); err != nil {
		return nil, err
	}

	now := time.Now()
	if err = r.CheckIn(now, reservation.CheckInDate, reservation.CheckOutDate); err != nil {
		return nil, err
	}

	wallet, err := r.walletRepo.GetWalletByUserId(userId)
	if err != nil {
		return nil, err
	}

	nights := reservation.CheckOutDate.Sub(reservation.CheckInDate).Hours() / 24
	if nights > 30 {
		return nil, errResp.ErrFailedCheckInMoreThan30Days
	}
	totalPrice := house.PricePerNight * float64(nights)

	if wallet.Balance < totalPrice {
		return nil, errResp.ErrInsufficientBalance
	}
	reservation.TotalPrice = totalPrice

	res, err := r.reservationRepo.CreateReservation(&reservation, hostId, pickup)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *reservationUseCaseImpl) CheckIn(now time.Time, checkInTime time.Time, checkOutTime time.Time) error {
	if checkInTime.Before(now) {
		return errResp.ErrFailedCheckInPast
	}

	if checkInTime.Equal(checkOutTime) {
		return errResp.ErrFailedCheckInWithSameCheckOut
	}

	return nil
}
