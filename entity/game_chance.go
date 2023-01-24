package entity

type GameChance struct {
	ID      uint `json:"id"`
	UserID  uint `json:"user_id"`
	Chances int  `json:"chance" gorm:"default:2"`
}

func (GameChance) TableName() string {
	return "games_chances"
}
