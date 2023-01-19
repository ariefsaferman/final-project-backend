package repository

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	TopUp(user entity.User) (string, error)
}

type transactionRepositoryImpl struct {
	db *gorm.DB
}

type TransactionRConfig struct {
	DB *gorm.DB
}

func NewTransactionRepository(config *TransactionRConfig) TransactionRepository {
	return &transactionRepositoryImpl{
		db: config.DB,
	}
}

func (r *transactionRepositoryImpl) TopUp(user entity.User) (string, error) {
	message := "successfuly top up wallet"
	var wallet entity.Wallet

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := r.db.Model(&wallet).Where("id = ?", user.Wallet.ID).Update("balance", gorm.Expr("balance + ?", user.Wallet.Balance)).Error; err != nil {
			return errors.ErrFailedToUpdateWallet
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	return message, nil
}
