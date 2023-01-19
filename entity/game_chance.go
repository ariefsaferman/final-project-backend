package entity

import "gorm.io/gorm"

type GameChance struct {
	ID uint 
	UserID uint 
	Chances int 
	gorm.Model 
}