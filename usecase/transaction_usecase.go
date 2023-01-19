package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/repository"
)

type TransactionUseCase interface {
	TopUp(req dto.TopUpRequest, UserID int) (*entity.Transaction, error)
}

type transactionUsecaseImpl struct {
	userRepo        repository.UserRepository
	transactionRepo repository.TransactionRepository
	walletRepo repository.WalletRepository
	sourceOfFundRepo repository.SourceOfFundRepository
}

type TransactionUConfig struct {
	UserRepo        repository.UserRepository
	TransactionRepo repository.TransactionRepository
	WalletRepo repository.WalletRepository
	SourceOfFundRepo repository.SourceOfFundRepository
}

func NewTransactionUsecase(config *TransactionUConfig) TransactionUseCase {
	return &transactionUsecaseImpl{
		userRepo:        config.UserRepo,
		transactionRepo: config.TransactionRepo,
		walletRepo: config.WalletRepo,
		sourceOfFundRepo: config.SourceOfFundRepo,
	}
}

func (u *transactionUsecaseImpl) TopUp(req dto.TopUpRequest, userID int) (*entity.Transaction, error) {
	req.SourceOfFundId = 1
	_, err := u.sourceOfFundRepo.FindByID(req.SourceOfFundId)
	if err != nil {
		return nil, err
	}

	transaction := req.ReqToTransaction(userID)
	res, err := u.transactionRepo.TopUp(transaction)
	if err != nil {
		return nil, err
	}

	return res, nil
}
