package repository

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"gorm.io/gorm"
)

type PickUpRepository interface {
	CreatePickUp(tx *gorm.DB, req entity.PickUp) error
}

type pickUpRepositoryImpl struct {
	db *gorm.DB
}

type PickUpRConfig struct {
	DB *gorm.DB
}

func NewPickUpRepository(config *PickUpRConfig) PickUpRepository {
	return &pickUpRepositoryImpl{
		db: config.DB,
	}
}

func (r *pickUpRepositoryImpl) CreatePickUp(tx *gorm.DB, req entity.PickUp) error {
	err := tx.Create(&req).Error
	if err != nil {
		return errResp.ErrFailedToCreatePickUp
	}
	return nil
}
