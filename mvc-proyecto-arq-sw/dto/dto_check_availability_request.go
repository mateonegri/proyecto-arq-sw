package dto

type CheckRoomDto struct {
	//StartDate datatypes.Date `json:"start_date"`
	//EndDate   datatypes.Date `json:"end_date"`

	/*StartDay   int `json:"start_day"`
	StartMonth int `json:"start_month"`
	StartYear  int `json:"start_year"`

	EndDay   int `json:"end_day"`
	EndMonth int `json:"end_month"`
	EndYear  int `json:"end_year"`*/

	StartDate int `json:"start_date"`
	EndDate   int `json:"end_date"`

	HotelId int `json:"booked_hotel_id"`
}

type CheckRoomsDto []CheckRoomDto
