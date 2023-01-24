package repository

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	TopUp(req entity.Transaction) (*entity.Transaction, error)
	TopUpWithReward(tx *gorm.DB, req entity.Transaction) (*entity.Transaction, error)
}

type transactionRepositoryImpl struct {
	db         *gorm.DB
	walletRepo WalletRepository
}

type TransactionRConfig struct {
	DB               *gorm.DB
	WalletRepository WalletRepository
}

func NewTransactionRepository(config *TransactionRConfig) TransactionRepository {
	return &transactionRepositoryImpl{
		db:         config.DB,
		walletRepo: config.WalletRepository,
	}
}

func (r *transactionRepositoryImpl) TopUp(req entity.Transaction) (*entity.Transaction, error) {

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.walletRepo.TopUp(tx, req.UserID, req.Balance); err != nil {
			return err
		}

		if err := tx.Create(&req).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &req, nil
}

func (r *transactionRepositoryImpl) TopUpWithReward(tx *gorm.DB, req entity.Transaction) (*entity.Transaction, error) {

	err := tx.Create(&req).Error
	if err != nil {
		return nil, err
	}

	return &req, nil
}
