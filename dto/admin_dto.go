package dto

import "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"

type AdminRegisterRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AdminRegisterResponse struct {
	ID     uint   `json:"id"`
	Email  string `json:"email"`
	RoleID int    `json:"role_id"`
}

func (r *AdminRegisterRequest) ReqToAdmin() entity.Admin {
	return entity.Admin{
		Email:    r.Email,
		Password: r.Password,
	}
}

func (r *AdminRegisterResponse) AdminToRes(admin entity.Admin) {
	r.ID = admin.ID
	r.Email = admin.Email
	r.RoleID = admin.RoleID
}