package response

import "github.com/gin-gonic/gin"

func SendSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{
		"data": data,
	})
}

func SendError(c *gin.Context, statusCode int, errCode string, message string) {
	c.JSON(statusCode, gin.H{
		"code": errCode,
		"message": message,
	})
}