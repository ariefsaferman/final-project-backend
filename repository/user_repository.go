package repository

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmailAndRole(email string, id int) (*entity.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

type UserRConfig struct {
	DB *gorm.DB
}

func NewUserRepository(config *UserRConfig) UserRepository {
	return &userRepositoryImpl{
		db: config.DB,
	}
}

func (r *userRepositoryImpl) FindByEmailAndRole(email string, id int) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).Where("role_id = ?", id).First(&user).Error
	if err != nil {
		return nil, errors.ErrUserNotFound
	}
	return &user, nil
}
