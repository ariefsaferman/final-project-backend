package entity

type GameChance struct {
	ID      uint `json:"id"`
	UserID  uint `json:"user_id"`
	Chances int  `json:"chance" default:"0"`
}

func (GameChance) TableName() string {
	return "games_chances"
}
