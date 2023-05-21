package model

type Booking struct {
	Id         int `gorm:"primaryKey"`
	StartDay   int `gorm:"not null"`
	StartMonth int `gorm:"not null"`
	StartYear  int `gorm:"not null"`

	EndDay   int `gorm:"not null"`
	EndMonth int `gorm:"not null"`
	EndYear  int `gorm:"not null"`

	User   User `gorm:"foreignkey:UserId"`
	UserId int

	Hotel   Hotel `gorm:"foreignkey:HotelId"`
	HotelId int
}

type Bookings []Booking
