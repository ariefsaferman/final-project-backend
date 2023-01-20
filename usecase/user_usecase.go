package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/constant"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/repository"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/auth"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
)

type UserUsecase interface {
	Login(req dto.LoginRequest, id int) (*dto.LoginResponse, error)
	AdminLogin(req dto.LoginRequest, id int) (*dto.LoginResponse, error)
	AdminRegister(req dto.AdminRegisterRequest) (*dto.AdminRegisterResponse, error)
	Register(req dto.RegisterRequest) (*dto.RegisterResponse, error)
	GetProfile(id int) (*entity.User, error)
	UpdateProfile(req dto.UpdateRequest, id int) (string, error)
	UpdateRole(id int) (string, error)
}

type userUsecaseImpl struct {
	userRepo      repository.UserRepository
	bcryptUseCase auth.AuthUtil
}

type UserUConfig struct {
	UserRepo      repository.UserRepository
	BcryptUseCase auth.AuthUtil
}

func NewUserUsecase(config *UserUConfig) UserUsecase {
	return &userUsecaseImpl{
		userRepo:      config.UserRepo,
		bcryptUseCase: config.BcryptUseCase,
	}
}

func (u *userUsecaseImpl) Register(req dto.RegisterRequest) (*dto.RegisterResponse, error) {
	userRegister := req.ReqToUser()
	userRegister.Password = u.bcryptUseCase.HashAndSalt(req.Password)
	userRegister.RoleID = constant.USER_ID
	if len(userRegister.Password) == 0 {
		return nil, errors.ErrFailedToHashPassword
	}

	user, err := u.userRepo.Register(userRegister)
	if err != nil {
		return nil, err
	}

	var res dto.RegisterResponse
	res.UserToRes(*user)

	return &res, nil
}

func (u *userUsecaseImpl) AdminRegister(req dto.AdminRegisterRequest) (*dto.AdminRegisterResponse, error) {
	adminRegister := req.ReqToAdmin()
	adminRegister.Password = u.bcryptUseCase.HashAndSalt(req.Password)
	adminRegister.RoleID = constant.ADMIN_ID
	if len(adminRegister.Password) == 0 {
		return nil, errors.ErrFailedToHashPassword
	}

	admin, err := u.userRepo.RegisterAdmin(adminRegister)
	if err != nil {
		return nil, err
	}

	var res dto.AdminRegisterResponse
	res.AdminToRes(*admin)

	return &res, nil
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

	user, err := u.userRepo.FindAdmin(req.Email)
	if err != nil {
		return nil, err
	}

	if ok := u.bcryptUseCase.ComparePasswords(user.Password, req.Password); !ok {
		return nil, errors.ErrWrongPassword
	}

	adminUser := entity.User{
		Email:    user.Email,
		Password: user.Password,
		RoleID:   uint(user.RoleID),
	}
	accessToken := u.bcryptUseCase.GenerateToken(adminUser)
	if accessToken.AccessToken == "" {
		return nil, errors.ErrFailedToGenerateToken
	}

	return &accessToken, nil
}

func (u *userUsecaseImpl) GetProfile(id int) (*entity.User, error) {
	user, err := u.userRepo.GetProfile(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecaseImpl) UpdateProfile(req dto.UpdateRequest, id int) (string, error) {
	user, err := u.userRepo.GetProfile(id)
	if err != nil {
		return "", err
	}

	user.FullName = req.FullName
	user.Password = u.bcryptUseCase.HashAndSalt(req.Password)
	user.Address = req.Address

	msg, err := u.userRepo.UpdateProfile(*user)
	if err != nil {
		return "", err
	}

	return msg, nil
}

func (u *userUsecaseImpl) UpdateRole(id int) (string, error) {
	user, err := u.userRepo.GetProfile(id)
	if err != nil {
		return "", err
	}

	user.RoleID = constant.HOST_ID

	msg, err := u.userRepo.UpdateRole(*user)
	if err != nil {
		return "", err
	}

	return msg, nil
}
