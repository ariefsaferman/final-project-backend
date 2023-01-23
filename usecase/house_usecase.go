package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/repository"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/media"
)

type HouseUseCase interface {
	CreateHouse(req dto.CreateHouseRequest, userID uint) (*entity.House, error)
	ViewHouseListHost(userId uint) ([]entity.House, error)
	ViewHouseList(query *entity.HouseParams) ([]entity.House, int64, int, error)
}

type houseUsecaseImpl struct {
	houseRepo repository.HouseRepository
}

type HouseUConfig struct {
	HouseRepo repository.HouseRepository
}

func NewHouseUsecase(config *HouseUConfig) HouseUseCase {
	return &houseUsecaseImpl{
		houseRepo: config.HouseRepo,
	}
}

func (u *houseUsecaseImpl) CreateHouse(req dto.CreateHouseRequest, userID uint) (*entity.House, error) {
	var house entity.House
	house.Name = req.Name
	house.UserID = userID
	house.PricePerNight = req.PricePerNight
	house.Description = req.Description
	house.CityID = req.CityID
	house.MaxGuests = req.MaxGuests

	for _, photo := range req.HousePhoto {
		file, err := photo.Open()
		if err != nil {
			return nil, errResp.ErrOpenFileHeader
		}
		url, err2 := media.ImageUploadHelper(file)
		if err2 != nil {
			return nil, errResp.ErrUploadImage
		}
		house.HousePhoto = append(house.HousePhoto, entity.HousePhoto{PhotoURL: url})
		file.Close()
	}

	res, err := u.houseRepo.CreateHouse(house)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *houseUsecaseImpl) ViewHouseListHost(userId uint) ([]entity.House, error) {
	res, err := u.houseRepo.ViewHouseListHost(userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *houseUsecaseImpl) ViewHouseList(query *entity.HouseParams) ([]entity.House, int64, int, error) {
	res, totalRows, totalPages, err := u.houseRepo.ViewHouseList(query)
	if err != nil {
		return nil, 0, 0, err
	}

	return res, totalRows, totalPages, nil
}
