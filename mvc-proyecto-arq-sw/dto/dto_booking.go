package dto

type BookingDto struct {
	Id int `json:"booking_id"`

	//StartDate datatypes.Date `json:"start_date"`
	//EndDate   datatypes.Date `json:"end_date"`

	UserId int `json:"user_booked_id"`

	HotelId int `json:"booked_hotel_id"`

	StartDate int `json:"start_date"`
	EndDate   int `json:"end_date"`

	/*StartDay   int `json:"start_day"`
	StartMonth int `json:"start_month"`
	StartYear  int `json:"start_year"`

	//Nights int `json:"nights"`

	EndDay   int `json:"end_day"`
	EndMonth int `json:"end_month"`
	EndYear  int `json:"end_year"`*/
}

type BookingsDto []BookingDto
