package errors

import "errors"

const (
	ErrCodeBadRequest = "BAD_REQUEST"
	ErrCodeInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrCodeUnauthorized = "UNAUTHORIZED"
)

var (
	ErrInvalidBody = errors.New("invalid body request")
	ErrWrongPassword = errors.New("password mismatch")
	ErrFailedToGenerateToken = errors.New("failed to generate token")
	ErrUserNotFound = errors.New("user not found")
	ErrInternalServerError = errors.New("internal server error")
	ErrBadHeader = errors.New("bad header")
	ErrIncorrectFormatHeader = errors.New("incorrect format header")
	ErrUserAlreadyExist = errors.New("user already exist")
	ErrFailedToHashPassword = errors.New("failed to hash password")
	ErrFailedToRegister = errors.New("failed to register")
	ErrFailedToCreateWallet = errors.New("failed to create wallet")
	ErrFailedToUpdateProfile = errors.New("failed to update profile")
	ErrFailedToUpdateRole = errors.New("failed to update role")
	ErrUnauthorized = errors.New("unauthorized")
	ErrFailedToUpdateWallet = errors.New("failed to update wallet")
	ErrInvalidLessAmount = errors.New("amount cannot be less than 10.000")
	ErrInvalidMoreAmount = errors.New("amount cannot be greater than 100.000.000")
	ErrWalletNotFound = errors.New("wallet not found")
	ErrSourceofFundNotFound = errors.New("source of fund not found")
	ErrCreateGameChance = errors.New("failed to create game chance")
	ErrFailedToCreateHouse = errors.New("failed to create house")
)