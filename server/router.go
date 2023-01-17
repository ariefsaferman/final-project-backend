package server

import "github.com/gin-gonic/gin"

type RouterConfig struct {}

func NewRouter (cfg *RouterConfig) *gin.Engine {
	r := gin.Default()
	return r
}