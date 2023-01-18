package repository

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
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

func (r *userRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}