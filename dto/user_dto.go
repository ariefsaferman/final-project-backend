package dto

import "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	FullName string `json:"full_name" validate:"required,min=3"`
	Address  string `json:"address" validate:"required"`
	CityID   uint   `json:"city_id" validate:"required"`
}

type UpdateRequest  struct {
	FullName string `json:"full_name" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
	Address  string `json:"address" validate:"required"`
}

type UpdateRoleRequest struct {
	RoleID uint `json:"role_id" validate:"required"`
}

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Address string `json:"address"`
	CityID uint `json:"city_id"`
	RoleID uint `json:"role_id"`
	Wallet entity.Wallet `json:"wallet"`
}

func (r *RegisterRequest) ReqToUser() entity.User {
	return entity.User{
		Email:    r.Email,
		Password: r.Password,
		FullName: r.FullName,
		Address:  r.Address,
		CityID:   r.CityID,
	}
}

func (r *RegisterResponse) UserToRes(user entity.User) {
	r.ID = user.ID
	r.Email = user.Email
	r.FullName = user.FullName
	r.Address = user.Address
	r.CityID = user.CityID
	r.RoleID = user.RoleID
	r.Wallet = user.Wallet
}
