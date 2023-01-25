package entity

import "gorm.io/gorm"

type PickupStatus struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	Status     string `json:"status" validate:"required"`
	gorm.Model `json:"-"`
}
