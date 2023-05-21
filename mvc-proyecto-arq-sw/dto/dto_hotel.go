package dto

type HotelDto struct {
	Id               int    `json:"id"`
	HotelName        string `json:"hotel_name"`
	HotelDescription string `json:"hotel_description"`
	Rooms            int    `json:"hotel_rooms"`
	AvailableRooms   int    `json:"available_rooms"`
	Address          string `json:"hotel_address"`
	ImageURL         string `json:"hotel_image_url"`
}

type HotelsDto []HotelDto
