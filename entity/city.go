package entity

import "gorm.io/gorm"

type City struct {
	ID         uint   `json:"id" gorm:"primaryKey" `
	Name       string `json:"name" validate:"required"`
	gorm.Model `json:"-"`
}
