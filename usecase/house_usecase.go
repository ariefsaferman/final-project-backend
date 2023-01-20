package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/repository"
)

type HouseUseCase interface {
	CreateHouse(req dto.CreateHouseRequest, userID uint) (*entity.House, error)
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

	res, err := u.houseRepo.CreateHouse(house)
	if err != nil {
		return nil, err
	}

	return res, nil
}
