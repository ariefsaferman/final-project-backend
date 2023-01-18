package handler

import (
	"errors"
	"net/http"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {
	var req dto.LoginRequest
	err := c.ShouldBindJSON(&req) 
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidBody.Error())
		return
	}

	res, err := h.userUsecase.Login(req)
	if err != nil {
		if errors.Is(err, errResp.ErrUserNotFound) || errors.Is(err, errResp.ErrWrongPassword){
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return 
	}
	response.SendSuccess(c, http.StatusOK, res)
}