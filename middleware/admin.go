package middleware

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	roleID := c.GetInt("roleID")
	if roleID != 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    errors.ErrCodeUnauthorized,
			"message": "you are not admin",
		})
		return
	}
}