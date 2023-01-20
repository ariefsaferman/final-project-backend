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

	walletRepo := repository.NewWalletRepository(&repository.WalletRConfig{
		DB: db.Get(),
	})

	gameRepo := repository.NewGameRepository(&repository.GameRConfig{
		DB: db.Get(),
	})

	userRepo := repository.NewUserRepository(&repository.UserRConfig{
		DB:               db.Get(),
		WalletRepository: walletRepo,
		GameRepository:   gameRepo,
	})

	sourceOfFundRepo := repository.NewSourceOfFundRepository(&repository.SourceOfFundRConfig{
		DB: db.Get(),
	})

	transactionRepo := repository.NewTransactionRepository(&repository.TransactionRConfig{
		DB:               db.Get(),
		WalletRepository: walletRepo,
	})

	userUsecase := usecase.NewUserUsecase(&usecase.UserUConfig{
		UserRepo:      userRepo,
		BcryptUseCase: auth.AuthUtilImpl{},
	})

	transactionUsecase := usecase.NewTransactionUsecase(&usecase.TransactionUConfig{
		UserRepo:         userRepo,
		TransactionRepo:  transactionRepo,
		WalletRepo:       walletRepo,
		SourceOfFundRepo: sourceOfFundRepo,
	})

	return NewRouter(&RouterConfig{
		UserUseCase:        userUsecase,
		TransactionUseCase: transactionUsecase,
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
