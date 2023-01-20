package server

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/handler"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/middleware"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/usecase"
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
	r.POST("/upload-house", middleware.Authenticated, middleware.AuthorizeHost, h.CreateHouse)
	return r
}
