package handler

import (
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
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
