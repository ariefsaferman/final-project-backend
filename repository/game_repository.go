package repository

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"gorm.io/gorm"
)

type GameRepository interface {
	CreateChance(tx *gorm.DB, userID uint) (*entity.GameChance, error)
	PlayGame(req *entity.GameChance, amount float64) (*dto.PlayGameResponse, error)
	GetChance(userID uint) (*entity.GameChance, error)
	UpdateChance(req *entity.GameChance) error
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

func (r *gameRepositoryImpl) GetChance(userID uint) (*entity.GameChance, error) {
	req := &entity.GameChance{}
	err := r.db.Where("user_id = ?", userID).First(&req).Error
	if err != nil {
		return nil, errResp.ErrGameChanceNotFound
	}

	return req, nil
}

func (r *gameRepositoryImpl) UpdateChance(req *entity.GameChance) error {
	err := r.db.Save(&req).Error
	if err != nil {
		return errResp.ErrUpdateGameChance
	}

	return nil
}

func (r *gameRepositoryImpl) PlayGame(req *entity.GameChance, amount float64) (*dto.PlayGameResponse, error) {

	res := dto.PlayGameResponse{}

	err := r.db.Transaction(func(tx *gorm.DB) error {

		//update chance
		req.Chances = req.Chances - 1
		req.NumberOfPlay = req.NumberOfPlay + 1
		if req.Chances < 0 {
			return errResp.ErrGameChanceNotEnough
		}

		//check number of play
		if req.NumberOfPlay%10 == 0 {
			req.Chances = req.Chances + 1
			res.Message = "Congratulations you got 1 chance to play again!"
		}

		err := tx.Save(&req).Error
		if err != nil {
			return errResp.ErrUpdateGameChance
		}

		//update wallet
		err = r.walletRepo.Reward(tx, req.UserID, amount)
		if err != nil {
			return err
		}

		//create transaction
		transaction := entity.Transaction{
			UserID:          req.UserID,
			Balance:         amount,
			SourceOfFundsId: 3,
		}
		_, err = r.transactionRepo.TopUpWithReward(tx, transaction)
		if err != nil {
			return err
		}

		res.Reward = amount
		return nil
	})

	if err != nil {
		return nil, err
	}
	return &res, nil
}
