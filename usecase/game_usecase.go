package usecase

import (
	"math/rand"
	"time"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/repository"
)

type GameUseCase interface {
	PlayGame(req *dto.PlayGameRequest, userId uint) (*dto.PlayGameResponse, error)
}

type gameUseCaseImpl struct {
	gameRepo   repository.GameRepository
	walletRepo repository.WalletRepository
}

type GameUConfig struct {
	GameRepo   repository.GameRepository
	WalletRepo repository.WalletRepository
}

func NewGameUseCase(config *GameUConfig) GameUseCase {
	return &gameUseCaseImpl{
		gameRepo:   config.GameRepo,
		walletRepo: config.WalletRepo,
	}
}

func (u *gameUseCaseImpl) PlayGame(req *dto.PlayGameRequest, userId uint) (*dto.PlayGameResponse, error) {
	min := 0
	max := 1000000
	gift := make([]int, 2)
	var err error

	//random number
	rand.Seed(time.Now().UnixNano())
	gift[0] = rand.Intn(max-min) + min
	gift[1] = rand.Intn(max-min) + min

	for i := 0; i < len(gift); i++ {
		if i == *req.SelectedInput {
			_, err = u.gameRepo.PlayGame(userId, float64(gift[i]))
			if err != nil {
				return nil, err
			}
			break
		}
	}

	res := &dto.PlayGameResponse{
		Amount: float64(gift[*req.SelectedInput]),
	}

	return res, nil
}
