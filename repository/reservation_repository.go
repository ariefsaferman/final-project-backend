package repository

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/constant"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"gorm.io/gorm"
)

type ReservationRepositry interface {
	CheckVacancy(houseId uint, reservation entity.Reservation) error
	CreateReservation(reservation *entity.Reservation, hostId uint, pickup *entity.PickUp) (*entity.Reservation, error)
}

type reservationRepositoryImpl struct {
	db              *gorm.DB
	transactionRepo TransactionRepository
	walletRepo      WalletRepository
	houseRepo       HouseRepository
	pickupRepo      PickUpRepository
}

type ReservationRConfig struct {
	DB              *gorm.DB
	TransactionRepo TransactionRepository
	WalletRepo      WalletRepository
	HouseRepo       HouseRepository
	PickUpRepo      PickUpRepository
}

func NewReservationRepository(config *ReservationRConfig) ReservationRepositry {
	return &reservationRepositoryImpl{
		db:              config.DB,
		transactionRepo: config.TransactionRepo,
		walletRepo:      config.WalletRepo,
		houseRepo:       config.HouseRepo,
		pickupRepo:      config.PickUpRepo,
	}
}

func (r *reservationRepositoryImpl) CreateReservation(reservation *entity.Reservation, hostId uint, pickup *entity.PickUp) (*entity.Reservation, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&reservation).Error; err != nil {
			return errResp.ErrFailedToCreateReservation
		}

		transaction := &entity.Transaction{
			UserID:          reservation.UserID,
			Balance:         reservation.TotalPrice,
			SourceOfFundsId: constant.DEBIT,
		}

		if err := r.transactionRepo.RentHouse(tx, transaction); err != nil {
			return err
		}

		if err := r.walletRepo.DeductBalance(tx, reservation.UserID, reservation.TotalPrice); err != nil {
			return err
		}

		if err := r.houseRepo.CheckIfRentOwnHouse(reservation.UserID, reservation.HouseID); err != nil {
			return err
		}

		if err := r.walletRepo.AddBalance(tx, hostId, reservation.TotalPrice); err != nil {
			return err
		}

		if pickup.UserID != 0 {
			pickup.ReservationID = reservation.ID
			if err := r.pickupRepo.CreatePickUp(tx, *pickup); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return reservation, nil
}

func (r *reservationRepositoryImpl) CheckVacancy(houseId uint, reservation entity.Reservation) error {
	err := r.db.Where("house_id = ?", houseId).Where("check_in_date <= ?", reservation.CheckInDate).Where("check_out_date >= ?", reservation.CheckInDate).First(&reservation).Error
	if err == nil {
		return errResp.ErrVacancyIsNotAvailable
	}

	if err := r.db.Where("house_id = ?", houseId).Where("check_in_date <= ?", reservation.CheckOutDate).Where("check_in_date >= ?", reservation.CheckInDate).First(&reservation).Error; err == nil {
		return errResp.ErrVacancyIsNotAvailable
	}

	return nil
}
