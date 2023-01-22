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
	db             *gorm.DB
	housePhotoRepo HousePhotoRepository
}

type HouseRConfig struct {
	DB             *gorm.DB
	HousePhotoRepo HousePhotoRepository
}

func NewHouseRepository(config *HouseRConfig) HouseRepository {
	return &houseRepositoryImpl{
		db:             config.DB,
		housePhotoRepo: config.HousePhotoRepo,
	}
}

func (r *houseRepositoryImpl) CreateHouse(req entity.House) (*entity.House, error) {

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&req).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {	
		return nil, errResp.ErrFailedToCreateHouse
	}
	return &req, nil
}
