package model

type Image struct {
	Id      int    `gorm:"primaryKey"`
	Url     string `gorm:"type:varchar(500);not null"`

	Hotel   Hotel `gorm:"foreignKey:HotelId"`
	HotelId int
}

type Images []Image