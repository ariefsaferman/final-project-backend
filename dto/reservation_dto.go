package dto

import (
	"time"

	"git.garena.com/sea-labs-id/batch-05/arief-saferman/house-booking/entity"
)

type ReservationRequest struct {
	HouseID       uint      `json:"house_id" validate:"required"`
	CheckInDate   time.Time `json:"check_in" validate:"required"`
	CheckOutDate  time.Time `json:"check_out" validate:"required"`
	PickupAddress string    `json:"pickup_address" validate:"required_with=CityID"`
	CityID        *uint     `json:"city_id" validate:"required_with=PickupAddress"`
}

func (r *ReservationRequest) ReqToReservation(userId uint) entity.Reservation {
	return entity.Reservation{
		UserID:       userId,
		HouseID:      r.HouseID,
		CheckInDate:  r.CheckInDate,
		CheckOutDate: r.CheckOutDate,
	}
}
