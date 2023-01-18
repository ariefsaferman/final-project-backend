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
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	FullName string `json:"full_name" binding:"required"`
	Address  string `json:"address" binding:"required"`
	CityID   uint   `json:"city_id" binding:"required"`
}

type RegisterResponse struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
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
}
