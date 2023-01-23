package usecase

import (
	"strings"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/repository"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/media"
)

type HouseUseCase interface {
	CreateHouse(req dto.CreateHouseRequest, userID uint) (*entity.House, error)
	ViewHouseListHost(userId uint, query *entity.HouseParams) ([]entity.House, int64, int, error)
	ViewHouseList(query *entity.HouseParams) ([]entity.House, int64, int, error)
	GetHouseDetail(id uint) (*entity.House, error)
	UpdateHouse(id uint, req dto.UpdateHouseRequest) (*entity.House, error)
}

type houseUsecaseImpl struct {
	houseRepo repository.HouseRepository
}

type HouseUConfig struct {
	HouseRepo repository.HouseRepository
}

func NewHouseUsecase(config *HouseUConfig) HouseUseCase {
	return &houseUsecaseImpl{
		houseRepo: config.HouseRepo,
	}
}

func (u *houseUsecaseImpl) CreateHouse(req dto.CreateHouseRequest, userID uint) (*entity.House, error) {
	var house entity.House
	house.Name = req.Name
	house.UserID = userID
	house.PricePerNight = req.PricePerNight
	house.Description = req.Description
	house.CityID = req.CityID
	house.MaxGuests = req.MaxGuests

	for _, photo := range req.HousePhoto {
		file, err := photo.Open()
		if err != nil {
			return nil, errResp.ErrOpenFileHeader
		}
		url, err2 := media.ImageUploadHelper(file)
		if err2 != nil {
			return nil, errResp.ErrUploadImage
		}
		house.HousePhoto = append(house.HousePhoto, entity.HousePhoto{PhotoURL: url})
		file.Close()
	}

	res, err := u.houseRepo.CreateHouse(house)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *houseUsecaseImpl) ViewHouseListHost(userId uint, query *entity.HouseParams) ([]entity.House, int64, int, error) {
	res, totalRows, totalPages, err := u.houseRepo.ViewHouseListHost(userId, query)
	if err != nil {
		return nil, 0, 0, err
	}

	return res, totalRows, totalPages, nil
}

func (u *houseUsecaseImpl) ViewHouseList(query *entity.HouseParams) ([]entity.House, int64, int, error) {
	res, totalRows, totalPages, err := u.houseRepo.ViewHouseList(query)
	if err != nil {
		return nil, 0, 0, err
	}

	return res, totalRows, totalPages, nil
}

func (u *houseUsecaseImpl) GetHouseDetail(id uint) (*entity.House, error) {
	res, err := u.houseRepo.GetHouseById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *houseUsecaseImpl) UpdateHouse(id uint, req dto.UpdateHouseRequest) (*entity.House, error) {
	house, err := u.houseRepo.GetHouseById(id)

	if err != nil {
		return nil, err
	}

	deletePhotoId := make(map[int]bool)
	for _, photo := range req.PhotoID {
		deletePhotoId[photo] = true
	}

	for idx, photo := range house.HousePhoto {
		if _, ok := deletePhotoId[int(photo.ID)]; ok {
			str := strings.Split(photo.PhotoURL, "/")
			publicId := strings.Split(str[len(str)-1], ".")[0]
			media.DeleteImageHelper(publicId)
			house.HousePhoto = append(house.HousePhoto[:idx], house.HousePhoto[idx+1:]...)
		}

	}

	house.Name = req.Name
	house.PricePerNight = req.PricePerNight
	house.Description = req.Description
	house.MaxGuests = req.MaxGuests
	for _, photo := range req.HousePhoto {
		file, err := photo.Open()
		if err != nil {
			return nil, errResp.ErrOpenFileHeader
		}
		url, err2 := media.ImageUploadHelper(file)
		if err2 != nil {
			return nil, errResp.ErrUploadImage
		}
		house.HousePhoto = append(house.HousePhoto, entity.HousePhoto{PhotoURL: url})
		file.Close()
	}

	res, err := u.houseRepo.UpdateHouse(*house)
	if err != nil {
		return nil, err
	}

	return res, nil
}
