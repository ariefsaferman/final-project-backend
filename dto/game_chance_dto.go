package dto

type GameChanceRequest struct {
	UserID  uint `json:"user_id"`
	Chances int  `json:"chance" default:"0"`
}
