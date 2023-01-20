package entity

type Admin struct {
	ID       uint   `json:"id" gorm:"primaryKey" `
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	RoleID   int    `json:"role_id" validate:"required"`
}

func (a *Admin) TableName() string {
	return "admin"
}
