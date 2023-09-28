package model

type Hotel struct {
	Id               int    `gorm:"primaryKey"`
	HotelName        string `gorm:"type:varchar(400);not null"`
	HotelDescription string `gorm:"type:varchar(600)"`
	Rooms            int    `gorm:"not null;default:5"`
	Address          string `gorm:"not null;type:varchar(500)"`
	ImageURL         string `gorm:"type:varchar(600)"`

	Amenities []*Amenitie `gorm:"many2many:hotels_amenities;"`
}

type Hotels []Hotel
