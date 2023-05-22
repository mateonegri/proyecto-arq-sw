package dto

type BookingDto struct {
	Id int `json:"booking_id"`

	StartDay   int `json:"start_day"`
	StartMonth int `json:"start_month"`
	StartYear  int `json:"start_year"`

	EndDay   int `json:"end_day"`
	EndMonth int `json:"end_month"`
	EndYear  int `json:"end_year"`

	UserId int `json:"user_booked_id"`

	HotelId int `json:"booked_hotel_id"`
}

type BookingsDto []BookingDto
