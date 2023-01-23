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
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	r := gin.Default()
	h := handler.NewHandler(&handler.Config{
		UserUsecase:        cfg.UserUseCase,
		TransactionUsecase: cfg.TransactionUseCase,
		HouseUsecase:       cfg.HouseUsecase,
	})

	r.NoRoute(func(c *gin.Context) {
		response.SendError(c, http.StatusNotFound, errResp.ErrCodePageNotFound, errResp.ErrPageNotFound.Error())
	})

	admin := r.Group("/admin")
	{
		admin.POST("/register", middleware.Authenticated, middleware.AuthorizeAdmin, h.RegisterAdmin)
		admin.POST("/login", h.AdminLogin)
	}

	r.POST("/login", h.Login)
	r.POST("/register", h.Register)
	r.GET("/profile", middleware.Authenticated, middleware.AuthorizeUserOrHost, h.GetProfile)
	r.PUT("/update-profile", middleware.Authenticated, middleware.AuthorizeUserOrHost, h.UpdateProfile)
	r.PUT("/update-role", middleware.Authenticated, middleware.AuthorizeUserOrHost, h.UpdateRole)
	r.POST("/top-up", middleware.Authenticated, middleware.AuthorizeUserOrHost, h.TopUp)
	r.PATCH("/houses/:id", middleware.Authenticated, middleware.AuthorizeHost, h.UpdateHouse)
	r.DELETE("/houses/:id", middleware.Authenticated, middleware.AuthorizeHost, h.DeleteHouse)
	r.GET("/houses", h.GetHouseList)
	r.GET("/houses/:id", h.GetHouseDetail)

	host := r.Group("/host")
	{
		host.POST("/upload-house", middleware.Authenticated, middleware.AuthorizeHost, h.CreateHouse)
		host.GET("/houses", middleware.Authenticated, middleware.AuthorizeHost, h.GetHouseListHost)
	}
	return r
}
