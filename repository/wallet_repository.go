package repository

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"gorm.io/gorm"
)

type WalletRepository interface {
	TopUp(tx *gorm.DB, walletID uint, amount float64) error
	CreateWallet(tx *gorm.DB, id uint) (*entity.Wallet, error)
}

type walletRepositoryImpl struct {
	db *gorm.DB
}

type WalletRConfig struct {
	DB *gorm.DB
}

func NewWalletRepository(config *WalletRConfig) WalletRepository {
	return &walletRepositoryImpl{
		db: config.DB,
	}
}

func (r *walletRepositoryImpl) CreateWallet(tx *gorm.DB, id uint) (*entity.Wallet, error){
	req := entity.Wallet{UserID: id}

	err := r.db.Create(&req).Error
	if err != nil {
		return nil, errors.ErrFailedToCreateWallet
	}
	return &req, nil
}

func (r *walletRepositoryImpl) TopUp(tx *gorm.DB, userID uint, amount float64) error {
	err := r.db.Model(&entity.Wallet{}).Where("user_id = ?", userID).Update("balance", gorm.Expr("balance + ?", amount))
	if err.Error != nil {
		return err.Error
	}

	if err.RowsAffected == 0 {
		return errors.ErrWalletNotFound
	}

	return nil
}