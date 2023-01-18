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
	Register(req dto.RegisterRequest) (string, error)
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

func (u *userUsecaseImpl) Register(req dto.RegisterRequest) (string, error) {
	userRegister := req.ReqToUser() 
	userRegister.Password = u.bcryptUseCase.HashAndSalt(req.Password)
	if len(userRegister.Password) == 0 {
		return "", errors.ErrFailedToHashPassword
	}

	msg, err := u.userRepo.Register(userRegister)
	if err != nil {
		return "", err
	}


	return msg, nil
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