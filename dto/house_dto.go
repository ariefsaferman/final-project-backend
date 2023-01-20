package dto

import (
	"mime/multipart"
)

type CreateHouseRequest struct {
	CityID        uint                  `form:"city_id" validate:"required"`
	Name          string                `form:"name" validate:"required"`
	PricePerNight float64               `form:"price_per_night" validate:"required"`
	Description   string                `form:"description" validate:"required"`
	MaxGuests     uint                  `form:"max_guest" validate:"required"`
	HousePhoto    *multipart.FileHeader `form:"house_photo"`
}

type HousePhotoRequest struct {
	PhotoUrl *multipart.FileHeader `form:"photo_url" validate:"required"`
}
