package entity

import "gorm.io/gorm"

type PickupStatus struct {
	ID uint
	Status string
	gorm.Model 
}