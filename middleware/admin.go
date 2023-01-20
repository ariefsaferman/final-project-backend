package middleware

import (
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/constant"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"github.com/gin-gonic/gin"
)

func AuthorizeAdmin(c *gin.Context) {
	roleID := c.GetInt("roleID")
	if roleID == constant.HOST_ID || roleID == constant.USER_ID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    errors.ErrCodeUnauthorized,
			"message": "you are not admin/superadmin",
		})
		return
	}
}

func AuthorizeUserOrHost(c *gin.Context) {
	roleID := c.GetInt("roleID")
	if roleID == constant.ADMIN_ID || roleID == constant.SUPER_ADMIN_ID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    errors.ErrCodeUnauthorized,
			"message": "you are not user/host",
		})
		return
	}
}

func AuthorizeHost(c *gin.Context) {
	roleID := c.GetInt("roleID")
	if roleID != constant.HOST_ID {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    errors.ErrCodeUnauthorized,
			"message": "you are not a host",
		})
		return
	}
}
