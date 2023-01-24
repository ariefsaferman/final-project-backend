package usecase

import (
	"strconv"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/repository"
)

type TransactionUseCase interface {
	TopUp(req dto.TopUpRequest, UserID int) (*dto.TopUpResponse, error)
}

type transactionUsecaseImpl struct {
	userRepo         repository.UserRepository
	transactionRepo  repository.TransactionRepository
	walletRepo       repository.WalletRepository
	sourceOfFundRepo repository.SourceOfFundRepository
	gameRepo         repository.GameRepository
}

type TransactionUConfig struct {
	UserRepo         repository.UserRepository
	TransactionRepo  repository.TransactionRepository
	WalletRepo       repository.WalletRepository
	SourceOfFundRepo repository.SourceOfFundRepository
	GameRepo         repository.GameRepository
}

func NewTransactionUsecase(config *TransactionUConfig) TransactionUseCase {
	return &transactionUsecaseImpl{
		userRepo:         config.UserRepo,
		transactionRepo:  config.TransactionRepo,
		walletRepo:       config.WalletRepo,
		sourceOfFundRepo: config.SourceOfFundRepo,
		gameRepo:         config.GameRepo,
	}
}

func (u *transactionUsecaseImpl) TopUp(req dto.TopUpRequest, userID int) (*dto.TopUpResponse, error) {
	transaction := req.ReqToTransaction(userID)
	transaction.SourceOfFundsId = 1
	_, err := u.sourceOfFundRepo.FindByID(transaction.SourceOfFundsId)
	if err != nil {
		return nil, err
	}

	res, err := u.transactionRepo.TopUp(transaction)
	transResp := dto.TopUpResponse{}
	transResp = transResp.TransactionToRes(*res)

	if int(req.Amount)%500000 == 0 {
		chance := int(req.Amount) / 500000
		gameChance, err := u.gameRepo.GetChance(uint(userID))
		if err != nil {
			return nil, err
		}

		gameChance.Chances += chance
		if err := u.gameRepo.UpdateChance(gameChance); err != nil {
			return nil, err
		}

		transResp.Message = `You got ` + strconv.Itoa(chance) + ` chances to play the game`
	}
	if err != nil {
		return nil, err
	}

	return &transResp, nil
}
