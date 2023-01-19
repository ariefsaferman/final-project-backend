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
)