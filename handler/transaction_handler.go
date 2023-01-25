package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func (h *Handler) TopUp(c *gin.Context) {
	var req dto.TopUpRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidBody.Error())
		return
	}

	v := validator.New()
	if errValidator := v.Struct(req); errValidator != nil {
		if strings.Contains(errValidator.Error(), "Field validation for 'Amount' failed on the 'gte' tag") {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidLessAmount.Error())
			return
		}

		if strings.Contains(errValidator.Error(), "Field validation for 'Amount' failed on the 'lte' tag") {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidMoreAmount.Error())
			return
		}
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errValidator.Error())
		return
	}

	userID := c.GetInt("userID")
	fmt.Println("userID: ", userID)
	res, err := h.transactionUsecase.TopUp(req, userID)
	if err != nil {
		if errors.Is(err, errResp.ErrSourceofFundNotFound) || errors.Is(err, errResp.ErrWalletNotFound) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}
	response.SendSuccess(c, http.StatusOK, res)
}

func (h *Handler) GetTransaction(c *gin.Context) {
	userID := c.GetInt("userID")
	res, err := h.transactionUsecase.GetTransaction(userID)
	if err != nil {
		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}
	response.SendSuccess(c, http.StatusOK, res)
}
