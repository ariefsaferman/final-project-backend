package handler

import "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/usecase"

type Handler struct {
	userUsecase usecase.UserUsecase
	transactionUsecase usecase.TransactionUseCase
}

type Config struct {
	UserUsecase usecase.UserUsecase
	TransactionUsecase usecase.TransactionUseCase
}

func NewHandler(config *Config) *Handler {
	return &Handler{
		userUsecase: config.UserUsecase,
		transactionUsecase: config.TransactionUsecase,
	}
}