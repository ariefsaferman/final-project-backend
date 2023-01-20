package entity

import "gorm.io/gorm"

type HousePhoto struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	HouseID    uint   `json:"house_id"`
	PhotoURL   string `json:"photo_url"`
	gorm.Model `json:"-"`
}

func (a *HousePhoto) TableName() string {
	return "houses_photos"
}
