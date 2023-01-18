package handler

import "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/usecase"

type Handler struct {
	userUsecase usecase.UserUsecase
}

type Config struct {
	UserUsecase usecase.UserUsecase
}

func NewHandler(config *Config) *Handler {
	return &Handler{
		userUsecase: config.UserUsecase,
	}
}