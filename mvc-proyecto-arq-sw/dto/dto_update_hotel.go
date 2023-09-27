package dto

type HandleHotelDto struct {
	Id               int    `json:"id"`
	HotelName        string `json:"hotel_name"`
	HotelDescription string `json:"hotel_description"`
	Rooms            int    `json:"hotel_rooms"`
	Address          string `json:"hotel_address"`
	ImageURL         string `json:"hotel_image_url"`

	Amenities   []string    `json:"amenities"`

	UserId int `json:"user_id"`
}

type HandleHotelsDto []HandleHotelDto
