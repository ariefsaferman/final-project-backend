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

func (h *Handler) PlayGame(c *gin.Context) {
	var req dto.PlayGameRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {

		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidBody.Error())
		return
	}

	v := validator.New()
	if errValidator := v.Struct(req); errValidator != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errValidator.Error())
		return
	}

	intUserId := c.GetInt("userID")
	res, err := h.gameUseCase.PlayGame(&req, uint(intUserId))
	if err != nil {
		if errors.Is(err, errResp.ErrGameChanceNotEnough) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrGameChanceNotEnough.Error())
			return
		}
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}
	response.SendSuccess(c, http.StatusOK, res)
}
