package entity

import (
	"mime/multipart"

	"gorm.io/gorm"
)

type House struct {
	ID            uint         `json:"id" gorm:"primaryKey" `
	Name          string       `json:"name" validate:"required"`
	UserID        uint         `json:"user_id" validate:"required"`
	PricePerNight float64      `json:"price_per_night" validate:"required"`
	Description   string       `json:"description" validate:"required"`
	CityID        uint         `json:"city_id" validate:"required"`
	MaxGuests     uint         `json:"max_guests" gorm:"column:max_guest"`
	HousePhoto    []HousePhoto `form:"house_photo" json:"house_photo"`
	City          City         `json:"city"`
	gorm.Model    `json:"-"`
}

type HousePhotoRequest struct {
	PhotoUrl *multipart.FileHeader `form:"photo_url" validate:"required"`
}

type HouseParams struct {
	Search string
	SortBy string
	Sort   string
	Limit  int
	Page   int
}

func NewHouseParams(search, sortBy, sort string, limit, page int) *HouseParams {
	return &HouseParams{
		Search: search,
		SortBy: func() string {
			if sortBy != "" {
				return sortBy
			}
			return "name"
		}(),
		Sort: func() string {
			if sort != "" {
				return sort
			}
			return "desc"
		}(),
		Limit: func() int {
			if limit > 0 {
				return limit
			}
			return 10
		}(),
		Page: func() int {
			if page > 1 {
				return page
			}
			return 1
		}(),
	}
}
