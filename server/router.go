package server

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/handler"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/usecase"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	UserUseCase usecase.UserUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	r := gin.Default()
	h := handler.NewHandler(&handler.Config{
		UserUsecase: cfg.UserUseCase,
	})

	admin := r.Group("/admin")
	{
		admin.POST("/login", h.AdminLogin)
	}

	r.POST("/login", h.Login)
	return r
}
