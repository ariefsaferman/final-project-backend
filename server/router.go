package server

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/handler"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/middleware"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/usecase"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/response"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUseCase        usecase.UserUsecase
	TransactionUseCase usecase.TransactionUseCase
	HouseUsecase       usecase.HouseUseCase
	GameUseCase        usecase.GameUseCase
	ReservationUseCase usecase.ReservationUseCase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	r := gin.Default()
	h := handler.NewHandler(&handler.Config{
		UserUsecase:        cfg.UserUseCase,
		TransactionUsecase: cfg.TransactionUseCase,
		HouseUsecase:       cfg.HouseUsecase,
		GameUseCase:        cfg.GameUseCase,
		ReservationUseCase: cfg.ReservationUseCase,
	})

	r.NoRoute(func(c *gin.Context) {
		response.SendError(c, http.StatusNotFound, errResp.ErrCodePageNotFound, errResp.ErrPageNotFound.Error())
	})

	admin := r.Group("/admin")
	{
		admin.POST("/register", middleware.Authenticated, middleware.AuthorizeAdmin, h.RegisterAdmin)
		admin.POST("/login", h.AdminLogin)
	}

	user := r.Group("/user", middleware.Authenticated, middleware.AuthorizeUserOrHost)
	{
		user.GET("/profile", h.GetProfile)
		user.PUT("/update-profile", h.UpdateProfile)
		user.PUT("/update-role", h.UpdateRole)

	}

	r.POST("/login", h.Login)
	r.POST("/register", h.Register)
	r.POST("/top-up", middleware.Authenticated, middleware.AuthorizeUserOrHost, h.TopUp)
	r.PATCH("/houses/:id", middleware.Authenticated, middleware.AuthorizeHost, h.UpdateHouse)
	r.DELETE("/houses/:id", middleware.Authenticated, middleware.AuthorizeHost, h.DeleteHouse)
	r.GET("/houses", h.GetHouseList)
	r.GET("/houses/:id", h.GetHouseDetail)
	r.POST("/play-game", middleware.Authenticated, middleware.AuthorizeUserOrHost, h.PlayGame)
	r.GET("/transaction", middleware.Authenticated, middleware.AuthorizeUserOrHost, h.GetTransaction)
	r.POST("/reservation", middleware.Authenticated, middleware.AuthorizeUserOrHost, h.CreateReservation)

	host := r.Group("/host")
	{
		host.POST("/upload-house", middleware.Authenticated, middleware.AuthorizeHost, h.CreateHouse)
		host.GET("/houses", middleware.Authenticated, middleware.AuthorizeHost, h.GetHouseListHost)
	}
	return r
}
