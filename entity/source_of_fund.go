package entity

import "gorm.io/gorm"

type SourceOfFund struct {
	ID uint
	Name string
	gorm.Model `json:"-"`
}