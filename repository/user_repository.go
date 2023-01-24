package repository

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/constant"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmailAndRole(email string, id int) (*entity.User, error)
	FindAdmin(email string) (*entity.Admin, error)
	Register(user entity.User) (*entity.User, error)
	RegisterAdmin(admin entity.Admin) (*entity.Admin, error)
	GetProfile(id int) (*entity.User, error)
	UpdateProfile(user entity.User) (string, error)
	UpdateRole(user entity.User) (string, error)
	GetUserByID(tx *gorm.DB, id uint) (*entity.User, error)
}

type userRepositoryImpl struct {
	db         *gorm.DB
	walletRepo WalletRepository
	gameRepo   GameRepository
}

type UserRConfig struct {
	DB               *gorm.DB
	WalletRepository WalletRepository
	GameRepository   GameRepository
}

func NewUserRepository(config *UserRConfig) UserRepository {
	return &userRepositoryImpl{
		db:         config.DB,
		walletRepo: config.WalletRepository,
		gameRepo:   config.GameRepository,
	}
}

func (r *userRepositoryImpl) Register(user entity.User) (*entity.User, error) {

	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return errors.ErrFailedToRegister
		}

		wallet, err := r.walletRepo.CreateWallet(tx, user.ID)
		if err != nil {
			return err
		}
		user.Wallet = *wallet

		gameChances, err := r.gameRepo.CreateChance(tx, user.ID)
		if err != nil {
			return err
		}
		user.GameChance = *gameChances

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepositoryImpl) RegisterAdmin(admin entity.Admin) (*entity.Admin, error) {
	err := r.db.Create(&admin).Error
	if err != nil {
		return nil, errors.ErrFailedToRegister
	}
	return &admin, nil
}

func (r *userRepositoryImpl) FindByEmailAndRole(email string, id int) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).Where("role_id != 1").First(&user).Error
	if err != nil {
		return nil, errors.ErrUserNotFound
	}
	return &user, nil
}

func (r *userRepositoryImpl) FindAdmin(email string) (*entity.Admin, error) {
	var user entity.Admin
	err := r.db.Where("email = ?", email).Where("role_id = ?", constant.ADMIN_ID).Or("role_id = ?", constant.SUPER_ADMIN_ID).First(&user).Error
	if err != nil {
		return nil, errors.ErrUserNotFound
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetProfile(id int) (*entity.User, error) {
	var user entity.User
	err := r.db.Preload("GameChance").Preload("Wallet").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, errors.ErrUserNotFound
	}
	return &user, nil
}

func (r *userRepositoryImpl) UpdateProfile(user entity.User) (string, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return "", errors.ErrFailedToUpdateProfile
	}
	message := "successfuly update profile"
	return message, nil
}

func (r *userRepositoryImpl) UpdateRole(user entity.User) (string, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return "", errors.ErrFailedToUpdateRole
	}
	message := "successfuly update role"

	return message, nil
}

func (r *userRepositoryImpl) GetUserByID(tx *gorm.DB, id uint) (*entity.User, error) {
	var user entity.User
	err := tx.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, errors.ErrUserNotFound
	}
	return &user, nil
}
