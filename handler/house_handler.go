package handler

import (
	"errors"
	"net/http"
	"strconv"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func (h *Handler) CreateHouse(c *gin.Context) {
	var req dto.CreateHouseRequest
	err := c.ShouldBind(&req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidBody.Error())
		return
	}

	v := validator.New()
	if errValidator := v.Struct(req); errValidator != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errValidator.Error())
		return
	}

	userID := c.GetInt("userID")
	res, err := h.houseUsecase.CreateHouse(req, uint(userID))
	if err != nil {
		if errors.Is(err, errResp.ErrFailedToCreateHouse) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}
	response.SendSuccess(c, http.StatusOK, res)
}

func (h *Handler) GetHouseListHost(c *gin.Context) {
	userID := c.GetInt("userID")
	intLimit, _ := strconv.Atoi(c.Query("limit"))
	intPage, _ := strconv.Atoi(c.Query("page"))
	params := entity.NewHouseParams(c.Query("search"), c.Query("sortBy"), c.Query("sort"), intLimit, intPage)

	res, totalRows, totalPages, err := h.houseUsecase.ViewHouseListHost(uint(userID), params)
	if err != nil {
		if errors.Is(err, errResp.ErrFailedToViewHouseList) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}
	response.SendSuccessWithPagination(c, http.StatusOK, res, totalRows, totalPages)
}

func (h *Handler) GetHouseList(c *gin.Context) {
	intLimit, _ := strconv.Atoi(c.Query("limit"))
	intPage, _ := strconv.Atoi(c.Query("page"))
	params := entity.NewHouseParams(c.Query("search"), c.Query("sortBy"), c.Query("sort"), intLimit, intPage)
	res, totalRows, totalPages, err := h.houseUsecase.ViewHouseList(params)
	if err != nil {
		if errors.Is(err, errResp.ErrFailedToViewHouseList) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}
	response.SendSuccessWithPagination(c, http.StatusOK, res, totalRows, totalPages)
}

func (h *Handler) GetHouseDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	res, err := h.houseUsecase.GetHouseDetail(uint(id))
	if err != nil {
		if errors.Is(err, errResp.ErrFailedToGetHouseDetail) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}
	response.SendSuccess(c, http.StatusOK, res)
}

func (h *Handler) UpdateHouse(c *gin.Context) {
	var req dto.UpdateHouseRequest
	err := c.ShouldBind(&req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidBody.Error())
		return
	}

	v := validator.New()
	if errValidator := v.Struct(req); errValidator != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errValidator.Error())
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	res, err := h.houseUsecase.UpdateHouse(uint(id), req)
	if err != nil {
		if errors.Is(err, errResp.ErrFailedToUpdateHouse) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}
	response.SendSuccess(c, http.StatusOK, res)
}

func (h *Handler) DeleteHouse(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.houseUsecase.DeleteHouse(uint(id))
	if err != nil {
		if errors.Is(err, errResp.ErrFailedToDeleteHouse) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}
	response.SendSuccess(c, http.StatusOK, "House Deleted")
}

