package dto

type TopUpRequest struct {
	Amount int `json:"amount" validate:"required, numeric, gte=10000, lte=100000000"`
}