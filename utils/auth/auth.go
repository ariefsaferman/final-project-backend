package auth

import (
	"time"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/config"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthUtil interface {
	HashAndSalt(pwd string) string
	ComparePasswords(hashedPwd string, plainPwd string) bool
	GenerateToken(req entity.User) dto.LoginResponse
}

type AuthUtilImpl struct{}

func NewAuthUtil() AuthUtilImpl {
	return AuthUtilImpl{}
}

func (a AuthUtilImpl) HashAndSalt(pwd string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)

	return string(hash)
}

func (a AuthUtilImpl) ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePlain := []byte(plainPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)

	return err == nil
}

type accessTokenClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	RoleID uint   `json:"role_id"`
	WalletID uint `json:"wallet_id"`
	jwt.RegisteredClaims
}

func (a AuthUtilImpl) GenerateToken(req entity.User) dto.LoginResponse {
	claims := accessTokenClaims{
		req.ID,
		req.Email,
		req.RoleID,
		req.Wallet.ID,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    config.AppName,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := token.SignedString([]byte(config.SecretKey))

	return dto.LoginResponse{AccessToken: signedToken}
}
