package handler

import "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/usecase"

type Handler struct {
	userUsecase usecase.UserUsecase
	transactionUsecase usecase.TransactionUseCase
	houseUsecase usecase.HouseUseCase
}

type Config struct {
	UserUsecase usecase.UserUsecase
	TransactionUsecase usecase.TransactionUseCase
	HouseUsecase usecase.HouseUseCase
}

func NewHandler(config *Config) *Handler {
	return &Handler{
		userUsecase: config.UserUsecase,
		transactionUsecase: config.TransactionUsecase,
		houseUsecase: config.HouseUsecase,
	}
}