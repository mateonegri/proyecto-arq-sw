package dto

type HandleHotelDto struct {
	Id               int    `json:"id"`
	HotelName        string `json:"hotel_name"`
	HotelDescription string `json:"hotel_description"`
	Rooms            int    `json:"hotel_rooms"`
	Address          string `json:"hotel_address"`
	Images      	 []string    `json:"images"`

	UserId int `json:"user_id"`
}

type HandleHotelsDto []HandleHotelDto
