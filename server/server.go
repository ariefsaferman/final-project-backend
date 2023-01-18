package server

import (
	"log"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/db"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/repository"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/usecase"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/auth"
	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {
	userRepo := repository.NewUserRepository(&repository.UserRConfig{
		DB: db.Get(), 
	})

	userUsecase := usecase.NewUserUsecase(&usecase.UserUConfig{
		UserRepo:     userRepo,
		BcryptUseCase: auth.AuthUtilImpl{},
	})

	return NewRouter(&RouterConfig{
		UserUseCase: userUsecase,
	})
}

func Init() {
	r := createRouter()
	err := r.Run()
	if err != nil {
		log.Println("error while running server", err)
		return
	}
}