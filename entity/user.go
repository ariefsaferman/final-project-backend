package entity

import "gorm.io/gorm"

type User struct {
	ID uint `json:"id"`
	Email string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Address string `json:"address"`
	CityID uint `json:"city_id"`
	RoleID uint `json:"role_id"`
	City City 
	Role Role
	Chance GameChance
	gorm.Model `json:"-"`
}