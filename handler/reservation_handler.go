package handler

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/dto"
	errResp "git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/errors"
	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func (h *Handler) CreateReservation(c *gin.Context) {
	var req *dto.ReservationRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(err)
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrInvalidBody.Error())
		return
	}

	v := validator.New()
	if errValidator := v.Struct(req); errValidator != nil {
		if strings.Contains(errValidator.Error(), "Field validation for 'CityID' failed on the 'required_with' tag") {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrCityIDRequired.Error())
			return
		}

		if strings.Contains(errValidator.Error(), "Field validation for 'PickupAddress' failed on the 'required_with' tag") {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errResp.ErrPickupAddressRequired.Error())
			return
		}
		
		response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, errValidator.Error())
		return
	}

	userID := c.GetInt("userID")
	res, err := h.reservationUseCase.CreateReservation(req, uint(userID))
	if err != nil {
		if errors.Is(err, errResp.ErrFailedToCreateReservation) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		if errors.Is(err, errResp.ErrRentOwnHouse) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		if errors.Is(err, errResp.ErrInsufficientBalance) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		if errors.Is(err, errResp.ErrHouseNotFound) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		if errors.Is(err, errResp.ErrVacancyIsNotAvailable) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		if errors.Is(err, errResp.ErrFailedCheckInPast) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		if errors.Is(err, errResp.ErrFailedCheckInWithSameCheckOut) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		if errors.Is(err, errResp.ErrFailedCheckInMoreThan30Days) {
			response.SendError(c, http.StatusBadRequest, errResp.ErrCodeBadRequest, err.Error())
			return
		}

		response.SendError(c, http.StatusInternalServerError, errResp.ErrCodeInternalServerError, errResp.ErrInternalServerError.Error())
		return
	}
	response.SendSuccess(c, http.StatusOK, res)

}
