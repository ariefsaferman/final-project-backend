package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/repository"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/auth"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
)

type UserUsecase interface {
	Login(req dto.LoginRequest, id int) (*dto.LoginResponse, error)
	AdminLogin(req dto.LoginRequest, id int) (*dto.LoginResponse, error)
} 

type userUsecaseImpl struct {
	userRepo     repository.UserRepository
	bcryptUseCase auth.AuthUtil
}

type UserUConfig struct {
	UserRepo     repository.UserRepository
	BcryptUseCase auth.AuthUtil
}

func NewUserUsecase(config *UserUConfig) UserUsecase {
	return &userUsecaseImpl{
		userRepo:     config.UserRepo,
		bcryptUseCase: config.BcryptUseCase,
	}
}

func (u *userUsecaseImpl) Login(req dto.LoginRequest, id int) (*dto.LoginResponse, error) {
	user, err := u.userRepo.FindByEmailAndRole(req.Email, id)
	if err != nil {
		return nil, err 
	}

	if ok := u.bcryptUseCase.ComparePasswords(user.Password, req.Password); !ok {
		return nil, errors.ErrWrongPassword
	}

	accessToken := u.bcryptUseCase.GenerateToken(*user) 
	if accessToken.AccessToken == "" {
		return nil, errors.ErrFailedToGenerateToken
	}

	return &accessToken, nil
}

func (u *userUsecaseImpl) AdminLogin(req dto.LoginRequest, id int) (*dto.LoginResponse, error) {
	user, err := u.userRepo.FindByEmailAndRole(req.Email, id)
	if err != nil {
		return nil, err 
	}

	if ok := u.bcryptUseCase.ComparePasswords(user.Password, req.Password); !ok {
		return nil, errors.ErrWrongPassword
	}

	accessToken := u.bcryptUseCase.GenerateToken(*user) 
	if accessToken.AccessToken == "" {
		return nil, errors.ErrFailedToGenerateToken
	}

	return &accessToken, nil
}