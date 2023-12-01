package dto

type AmenitieDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AmenitiesDto struct {
	Amenities []AmenitieDto `json:"amenities"`
}