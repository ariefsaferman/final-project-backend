package repository

import (
	"errors"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"gorm.io/gorm"
)

type SourceOfFundRepository interface {
	FindByID(id uint) (*entity.SourceOfFund, error)
}

type sourceOfFundRepositoryImpl struct {
	db *gorm.DB
}

type SourceOfFundRConfig struct {
	DB *gorm.DB
}

func NewSourceOfFundRepository(config *SourceOfFundRConfig) SourceOfFundRepository {
	return &sourceOfFundRepositoryImpl{
		db: config.DB,
	}
}

func (r *sourceOfFundRepositoryImpl) FindByID(id uint) (*entity.SourceOfFund, error) {
	var sourceOfFund entity.SourceOfFund
	err := r.db.First(&sourceOfFund, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errResp.ErrSourceofFundNotFound
		}
		return nil, err
	}

	return &sourceOfFund, nil
}
