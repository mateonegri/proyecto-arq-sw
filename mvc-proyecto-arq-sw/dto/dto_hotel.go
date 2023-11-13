package dto

type HotelDto struct {
	Id               int    `json:"id"`
	HotelName        string `json:"hotel_name"`
	HotelDescription string `json:"hotel_description"`
	Rooms            int    `json:"hotel_rooms"`
	Address          string `json:"hotel_address"`
	
	
	Amenities        []string `json:"amenities"`
	Images      []string    `json:"images"`
}

type HotelsDto []HotelDto
