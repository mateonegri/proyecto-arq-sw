package dto

type BookingDetailDto struct {
	StartDay   int `json:"start_day"`
	StartMonth int `json:"start_month"`
	StartYear  int `json:"start_year"`

	EndDay   int `json:"end_day"`
	EndMonth int `json:"end_month"`
	EndYear  int `json:"end_year"`

	UserId   int    `json:"user_booked_id"`
	Username string `json:"user_name"`

	HotelId   int    `json:"booked_hotel_id"`
	HotelName string `json:"hotel_name"`
	Address   string `json:"hotel_address"`
}

type BookingsDetailDto []BookingDetailDto
