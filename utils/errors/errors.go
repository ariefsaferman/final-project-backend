package errors

import "errors"

const (
	ErrCodeBadRequest          = "BAD_REQUEST"
	ErrCodePageNotFound        = "PAGE_NOT_FOUND"
	ErrCodeInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrCodeUnauthorized        = "UNAUTHORIZED"
)

var (
	ErrInvalidBody                   = errors.New("invalid body request")
	ErrPageNotFound                  = errors.New("page not found")
	ErrWrongPassword                 = errors.New("password mismatch")
	ErrFailedToGenerateToken         = errors.New("failed to generate token")
	ErrUserNotFound                  = errors.New("user not found")
	ErrInternalServerError           = errors.New("internal server error")
	ErrBadHeader                     = errors.New("bad header")
	ErrIncorrectFormatHeader         = errors.New("incorrect format header")
	ErrUserAlreadyExist              = errors.New("user already exist")
	ErrFailedToHashPassword          = errors.New("failed to hash password")
	ErrFailedToRegister              = errors.New("failed to register")
	ErrFailedToCreateWallet          = errors.New("failed to create wallet")
	ErrFailedToUpdateProfile         = errors.New("failed to update profile")
	ErrFailedToUpdateRole            = errors.New("failed to update role")
	ErrUnauthorized                  = errors.New("unauthorized")
	ErrFailedToUpdateWallet          = errors.New("failed to update wallet")
	ErrInvalidLessAmount             = errors.New("amount cannot be less than 10.000")
	ErrInvalidMoreAmount             = errors.New("amount cannot be greater than 100.000.000")
	ErrWalletNotFound                = errors.New("wallet not found")
	ErrSourceofFundNotFound          = errors.New("source of fund not found")
	ErrCreateGameChance              = errors.New("failed to create game chance")
	ErrFailedToCreateHouse           = errors.New("failed to create house")
	ErrFailedToCreateHousePhoto      = errors.New("failed to create house photo")
	ErrOpenFileHeader                = errors.New("failed to open file header")
	ErrUploadImage                   = errors.New("failed to upload image")
	ErrFailedToViewHouseList         = errors.New("failed to view house list")
	ErrFailedToGetHousePhoto         = errors.New("failed to get house photo")
	ErrFailedToGetHouseDetail        = errors.New("failed to get house detail")
	ErrFailedToUpdateHouse           = errors.New("failed to update house")
	ErrFailedToDeleteHousePhoto      = errors.New("failed to delete house photo")
	ErrFailedToDeleteHouse           = errors.New("failed to delete house")
	ErrUpdateGameChance              = errors.New("failed to update game chance")
	ErrGameChanceNotFound            = errors.New("game chance not with that user_id is not found")
	ErrGameChanceNotEnough           = errors.New("game chance not enough")
	ErrFailedToCreateReservation     = errors.New("failed to create reservation")
	ErrVacancyIsNotAvailable         = errors.New("house is not available on range date")
	ErrInsufficientBalance           = errors.New("insufficient balance")
	ErrFailToCreateRentTransaction   = errors.New("failed to create rent transaction")
	ErrRentOwnHouse                  = errors.New("you cannot rent your own house")
	ErrHouseNotFound                 = errors.New("house not found")
	ErrFailedCheckInPast             = errors.New("your reservation is in the past")
	ErrFailedCheckInWithSameCheckOut = errors.New("cannot check out with same date as check in")
	ErrFailedCheckInMoreThan30Days   = errors.New("cannot check in more than 30 days")
	ErrCityIDRequired                = errors.New("city_id is required")
	ErrPickupAddressRequired         = errors.New("pickup_address is required")
	ErrFailedToCreatePickUp          = errors.New("failed to create pickup")
)
