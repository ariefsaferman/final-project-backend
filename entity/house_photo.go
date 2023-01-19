package entity

import "gorm.io/gorm"

type HousePhoto struct { 
	ID uint
	HouseID uint
	PhotoURL string 
	gorm.Model 
}