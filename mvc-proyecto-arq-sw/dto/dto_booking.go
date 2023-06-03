package dto

type BookingDto struct {
	Id int `json:"booking_id"`

	//StartDate datatypes.Date `json:"start_date"`
	//EndDate   datatypes.Date `json:"end_date"`

	UserId int `json:"user_booked_id"`

	HotelId int `json:"booked_hotel_id"`

	StartDate int `json:"start_date"`
	EndDate   int `json:"end_date"`




}

type BookingsDto []BookingDto
