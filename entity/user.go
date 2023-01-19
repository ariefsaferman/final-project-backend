package entity

import "gorm.io/gorm"

type User struct {
	ID         uint       `json:"id"`
	Email      string     `json:"email"`
	Password   string     `json:"-"`
	FullName   string     `json:"full_name"`
	Address    string     `json:"address"`
	CityID     uint       `json:"city_id"`
	RoleID     uint       `json:"role_id"`
	Wallet 	   Wallet     `json:"wallet"`
	City       City       `json:"-"`
	Role       Role       `json:"-"`
	Chance     GameChance `json:"-"`
	gorm.Model `json:"-"`
}
