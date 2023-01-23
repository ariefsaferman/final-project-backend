package repository

import (
	"fmt"
	"math"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"gorm.io/gorm"
)

type HouseRepository interface {
	CreateHouse(req entity.House) (*entity.House, error)
	ViewHouseListHost(userId uint) ([]entity.House, error)
	ViewHouseList(query *entity.HouseParams) ([]entity.House, int64, int, error)
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

func (r *houseRepositoryImpl) ViewHouseListHost(userId uint) ([]entity.House, error) {
	var houses []entity.House
	err := r.db.Preload("HousePhoto").Where("user_id = ?", userId).Find(&houses).Error
	if err != nil {
		return nil, errResp.ErrFailedToViewHouseList
	}
	return houses, nil
}

func (r *houseRepositoryImpl) ViewHouseList(query *entity.HouseParams) ([]entity.House, int64, int, error) {
	var houses []entity.House
	var totalRows int64

	sort := fmt.Sprintf("%s %s", query.SortBy, query.Sort)
	offset := (query.Page - 1) * query.Limit

	db := r.db.Joins("inner join cities on cities.id = houses.city_id").Where(r.db.Where("houses.name ilike ?", "%"+query.Search+"%").Or("cities.name ilike ?", "%"+query.Search+"%")).Order(sort)
	db.Model(&houses).Count(&totalRows)
	totalPages := int(math.Ceil(float64(totalRows) / float64(query.Limit)))
	err := db.Preload("City").Preload("HousePhoto").Limit(query.Limit).Offset(offset).Find(&houses).Error
	if err != nil {
		return nil, 0, 0, errResp.ErrFailedToViewHouseList
	}

	return houses, totalRows, totalPages, nil
}
