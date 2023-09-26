package hotel

import (
	"mvc-proyecto-arq-sw/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetHotelById(id int) model.Hotel {
	var hotel model.Hotel

	Db.Where("id = ?", id).First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

func CheckHotelById(id int) bool {
	var hotel model.Hotel

	result := Db.Where("id = ?", id).First(&hotel)

	if result.Error != nil {
		return false
	}

	return true
}

func GetHotels() model.Hotels {
	var hotels model.Hotels
	Db.Find(&hotels)

	log.Debug("Hotels: ", hotels)

	return hotels
}

func InsertHotel(hotel model.Hotel) model.Hotel {
	result := Db.Create(&hotel)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
		hotel.Id = 0
	}
	log.Debug("Hotel Created: ", hotel.Id)
	return hotel
}

func UpdateHotelById(hotel model.Hotel) model.Hotel {

	result := Db.Model(&hotel).Where("id = ? ", hotel.Id).Updates(map[string]interface{}{"hotel_name": hotel.HotelName, "hotel_description": hotel.HotelDescription, "rooms": hotel.Rooms, "address": hotel.Address, "image_url": hotel.ImageURL})

	if result.Error != nil {
		log.Error("")
		hotel.Id = 0
	}

	log.Debug("Hotel Updated or Created: ", hotel.Id)

	return hotel
}

func DeleteHotel(hotel model.Hotel) (bool, error) {

	result := Db.Delete(&hotel, hotel.Id)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
