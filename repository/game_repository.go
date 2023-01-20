package repository

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"gorm.io/gorm"
)

type GameRepository interface {
	CreateChance(tx *gorm.DB, userID uint) (*entity.GameChance, error)
}

type gameRepositoryImpl struct {
	db *gorm.DB
}

type GameRConfig struct {
	DB *gorm.DB
}

func NewGameRepository(config *GameRConfig) GameRepository {
	return &gameRepositoryImpl{
		db: config.DB,
	}
}

func (r *gameRepositoryImpl) CreateChance(tx *gorm.DB, userID uint) (*entity.GameChance, error) {
	req := entity.GameChance{UserID: userID}

	err := tx.Create(&req).Error

	if err != nil {
		
		return nil, errResp.ErrCreateGameChance
	}

	return &req, nil
}
