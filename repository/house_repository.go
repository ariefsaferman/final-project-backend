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
	ViewHouseListHost(userId uint, query *entity.HouseParams) ([]entity.House, int64, int, error)
	ViewHouseList(query *entity.HouseParams) ([]entity.House, int64, int, error)
	GetHouseById(id uint) (*entity.House, error)
	UpdateHouse(req entity.House) (*entity.House, error)
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

func (r *houseRepositoryImpl) ViewHouseListHost(userId uint, query *entity.HouseParams) ([]entity.House, int64, int, error) {
	var houses []entity.House
	var totalRows int64
	sort := fmt.Sprintf("%s %s", query.SortBy, query.Sort)
	offset := (query.Page - 1) * query.Limit

	db := r.db.Joins("inner join cities on cities.id = houses.city_id").Where(r.db.Where("houses.name ilike ?", "%"+query.Search+"%").Or("cities.name ilike ?", "%"+query.Search+"%")).Order(sort)
	db.Model(&houses).Count(&totalRows)
	totalPages := int(math.Ceil(float64(totalRows) / float64(query.Limit)))
	err := db.Preload("City").Preload("HousePhoto").Where("user_id = ?", userId).Limit(query.Limit).Offset(offset).Find(&houses).Error
	if err != nil {
		return nil, 0, 0, errResp.ErrFailedToViewHouseList
	}

	return houses, totalRows, totalPages, nil
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

func (r *houseRepositoryImpl) GetHouseById(id uint) (*entity.House, error) {
	var house entity.House
	err := r.db.Preload("HousePhoto").Where("id = ?", id).First(&house).Error
	if err != nil {
		return nil, errResp.ErrFailedToGetHouseDetail
	}
	return &house, nil
}

func (r *houseRepositoryImpl) UpdateHouse(req entity.House) (*entity.House, error) {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.housePhotoRepo.DeleteHousePhotoByHouseId(tx, req.ID); err != nil {
			return err
		}

		if err := tx.Save(&req).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, errResp.ErrFailedToUpdateHouse
	}
	return &req, nil
}
