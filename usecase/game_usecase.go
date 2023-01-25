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
	gift := make([]int, 2)
	var err error
	var res *dto.PlayGameResponse

	//random number
	rand.Seed(time.Now().UnixNano())
	gift[0] = rand.Intn(10000-0) + 0
	gift[1] = rand.Intn(1000000-100000) + 100000
	rand.Shuffle(len(gift), func(i, j int) { gift[i], gift[j] = gift[j], gift[i] })

	//get game chance
	game, err := u.gameRepo.GetChance(userId)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(gift); i++ {
		if i == *req.SelectedInput {
			res, err = u.gameRepo.PlayGame(game, float64(gift[i]))
			if err != nil {
				return nil, err
			}
			break
		}
	}
	return res, nil
}
