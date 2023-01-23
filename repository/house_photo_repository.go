package repository

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"gorm.io/gorm"
)

type HousePhotoRepository interface {
	CreateHousePhoto(tx *gorm.DB, req *entity.HousePhoto) (*entity.HousePhoto, error)
	DeleteHousePhotoByHouseId(tx *gorm.DB, id uint) error
}

type housePhotoRepositoryImpl struct {
	db *gorm.DB
}

type HousePhotoRConfig struct {
	DB *gorm.DB
}

func NewHousePhotoRepository(config *HousePhotoRConfig) HousePhotoRepository {
	return &housePhotoRepositoryImpl{
		db: config.DB,
	}
}

func (r *housePhotoRepositoryImpl) CreateHousePhoto(tx *gorm.DB, req *entity.HousePhoto) (*entity.HousePhoto, error) {
	err := tx.Create(&req).Error
	if err != nil {
		return nil, errResp.ErrFailedToCreateHousePhoto
	}
	return req, nil
}

func (r *housePhotoRepositoryImpl) DeleteHousePhotoByHouseId(tx *gorm.DB, id uint) error {
	var req entity.HousePhoto
	err := tx.Where("house_id = ?", id).Delete(&req).Error
	if err != nil {
		return errResp.ErrFailedToDeleteHousePhoto
	}
	return nil
}
