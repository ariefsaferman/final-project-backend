package entity

import "gorm.io/gorm"

type Role struct {
	ID uint
	Name string 
	UserId uint 
	gorm.Model `json:"-"`
}