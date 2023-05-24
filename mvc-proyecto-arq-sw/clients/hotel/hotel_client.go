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
	var exist bool

	result := Db.Where("id = ?", id).First(&hotel)

	if result.Error != nil {
		exist = false
		return exist
	}

	exist = true
	return exist
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
	}
	log.Debug("Hotel Created: ", hotel.Id)
	return hotel
}

/*
func UpdateAvailableRooms() model.Hotel { //Preguntar esto!
	result := Db.UpdateColumn()
}
*/
