package entity

type Wallet struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint 
	Balance int `gorm:"default:0"`
}