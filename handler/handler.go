package handler

import "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/usecase"

type Handler struct {
	userUsecase        usecase.UserUsecase
	transactionUsecase usecase.TransactionUseCase
	houseUsecase       usecase.HouseUseCase
	gameUseCase        usecase.GameUseCase
}

type Config struct {
	UserUsecase        usecase.UserUsecase
	TransactionUsecase usecase.TransactionUseCase
	HouseUsecase       usecase.HouseUseCase
	GameUseCase        usecase.GameUseCase
}

func NewHandler(config *Config) *Handler {
	return &Handler{
		userUsecase:        config.UserUsecase,
		transactionUsecase: config.TransactionUsecase,
		houseUsecase:       config.HouseUsecase,
		gameUseCase:        config.GameUseCase,
	}
}
