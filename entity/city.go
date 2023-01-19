package entity

import "gorm.io/gorm"

type City struct {
	ID uint
	Name string 
	UserID uint 
	gorm.Model 
}