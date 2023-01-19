package entity

type Wallet struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint 
	Balance float64 `gorm:"default:0"`
}