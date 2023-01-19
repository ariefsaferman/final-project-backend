package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/repository"
)

type TransactionUseCase interface {
	TopUp(req dto.TopUpRequest, userId int, id int) (string, error)
}

type transactionUsecaseImpl struct {
	userRepo        repository.UserRepository
	transactionRepo repository.TransactionRepository
}

type TransactionUConfig struct {
	UserRepo        repository.UserRepository
	TransactionRepo repository.TransactionRepository
}

func NewTransactionUsecase(config *TransactionUConfig) TransactionUseCase {
	return &transactionUsecaseImpl{
		userRepo:        config.UserRepo,
		transactionRepo: config.TransactionRepo,
	}
}

func (u *transactionUsecaseImpl) TopUp(req dto.TopUpRequest, userId int, id int) (string, error) {
	user, err := u.userRepo.GetProfile(userId)
	if err != nil {
		return "", err
	}

	user.Wallet.Balance += req.Amount

	msg, err := u.transactionRepo.TopUp(*user)
	if err != nil {
		return "", err
	}

	return msg, nil
}
