package server

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/handler"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/middleware"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/usecase"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUseCase usecase.UserUsecase
	TransactionUseCase usecase.TransactionUseCase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	r := gin.Default()
	h := handler.NewHandler(&handler.Config{
		UserUsecase: cfg.UserUseCase,
		TransactionUsecase: cfg.TransactionUseCase,
	})

	admin := r.Group("/admin")
	{
		admin.POST("/register", h.RegisterAdmin)
		admin.POST("/login", h.AdminLogin)
	}

	r.POST("/login", h.Login)
	r.POST("/register", h.Register)
	r.GET("/profile", middleware.Authenticated, h.GetProfile)
	r.PUT("/update-profile", middleware.Authenticated, h.UpdateProfile)
	r.PUT("/update-role", middleware.Authenticated, h.UpdateRole)
	r.POST("/top-up", middleware.Authenticated, h.TopUp)
	return r
}