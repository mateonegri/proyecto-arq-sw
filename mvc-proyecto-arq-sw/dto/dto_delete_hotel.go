package dto

type DeleteHotelDto struct {
	HotelId int `json:"hotel_id"`
	UserId  int `json:"user_id"`
}
