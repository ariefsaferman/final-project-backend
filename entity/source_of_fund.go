package entity

import "gorm.io/gorm"

type SourceOfFund struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	gorm.Model `json:"-"`
}
