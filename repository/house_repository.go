package repository

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"gorm.io/gorm"
)

type HouseRepository interface {
	CreateHouse(req entity.House) (*entity.House, error)
}

type houseRepositoryImpl struct {
	db *gorm.DB
}

type HouseRConfig struct {
	DB *gorm.DB
}

func NewHouseRepository(config *HouseRConfig) HouseRepository {
	return &houseRepositoryImpl{
		db: config.DB,
	}
}

func (r *houseRepositoryImpl) CreateHouse(req entity.House) (*entity.House, error) {
	err := r.db.Create(&req).Error
	if err != nil {
		return nil, errResp.ErrFailedToCreateHouse
	}
	return &req, nil
}
