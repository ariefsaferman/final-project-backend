package middleware

import (
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/config"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authenticated(c *gin.Context) {
	tokenString, err := getJwtToken(c.GetHeader("Authorization"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    errors.ErrCodeUnauthorized,
			"message": err.Error(),
		})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    errors.ErrCodeUnauthorized,
					"message": "not a token",
				})
				return
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    errors.ErrCodeUnauthorized,
					"message": "token expired or not actived yet",
				})
				return
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    errors.ErrCodeUnauthorized,
					"message": "couldn't handle token",
				})
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    errors.ErrCodeUnauthorized,
			"message": "couldn't handle token",
		})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		intId := int(claims["user_id"].(float64))
		intRoleId := int(claims["role_id"].(float64))
		intWalletId := int(claims["wallet_id"].(float64))
		c.Set("userID", intId)
		c.Set("roleID", intRoleId)
		c.Set("walletID", intWalletId)
		c.Next()
		return
	}

	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"code":    errors.ErrCodeUnauthorized,
		"message": "couldn't handle token",
	})
}

func getJwtToken(header string) (string, error) {
	if header == "" {
		return "", errors.ErrBadHeader
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", errors.ErrBadHeader
	}

	return jwtToken[1], nil
}
