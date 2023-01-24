package repository

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"gorm.io/gorm"
)

type GameRepository interface {
	CreateChance(tx *gorm.DB, userID uint) (*entity.GameChance, error)
	PlayGame(userID uint, amount float64) (*entity.GameChance, error)
}

type gameRepositoryImpl struct {
	db              *gorm.DB
	walletRepo      WalletRepository
	transactionRepo TransactionRepository
}

type GameRConfig struct {
	DB              *gorm.DB
	WalletRepo      WalletRepository
	TransactionRepo TransactionRepository
}

func NewGameRepository(config *GameRConfig) GameRepository {
	return &gameRepositoryImpl{
		db:              config.DB,
		walletRepo:      config.WalletRepo,
		transactionRepo: config.TransactionRepo,
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

func (r *gameRepositoryImpl) PlayGame(userID uint, amount float64) (*entity.GameChance, error) {

	var req *entity.GameChance
	err := r.db.Transaction(func(tx *gorm.DB) error {
		//update chance
		err := tx.Where("user_id = ?", userID).First(&req).Error
		if err != nil {
			return errResp.ErrGameChanceNotFound
		}

		req.Chances = req.Chances - 1
		if req.Chances < 0 {
			return errResp.ErrGameChanceNotEnough
		}

		err = tx.Save(&req).Error
		if err != nil {
			return errResp.ErrUpdateGameChance
		}

		//update wallet
		err = r.walletRepo.Reward(tx, userID, amount)
		if err != nil {
			return err
		}

		//create transaction
		transaction := entity.Transaction{
			UserID:          userID,
			Balance:         amount,
			SourceOfFundsId: 3,
		}
		_, err = r.transactionRepo.TopUpWithReward(tx, transaction)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}
	return req, nil
}
