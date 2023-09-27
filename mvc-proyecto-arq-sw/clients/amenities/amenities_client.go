package amenities

import (
	"mvc-proyecto-arq-sw/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"errors"
	e "mvc-proyecto-arq-sw/utils/errors"
)

var Db *gorm.DB

func GetAmenitieById(id int) model.Amenitie {
	var amenitie model.Amenitie

	Db.Where("id = ?", id).First(&amenitie)
	log.Debug("Aminitie: ", amenitie)

	return amenitie
}

func GetAmenities() model.Amenities {
	var amenities model.Amenities
	Db.Find(&amenities)

	log.Debug("Amenities: ", amenities)

	return amenities
}

func InsertAmenitie(amenitie model.Amenitie) model.Amenitie {
	result := Db.Create(&amenitie)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Amenitie Created: ", amenitie.Id)
	return amenitie
}

func GetAmenitiesByHotelId(hotelId int) model.Amenities {
	var amenities model.Amenities

	Db.Table("amenities").
		Joins("JOIN hotels_amenities ON amenities.id = hotels_amenities.amenitie_id").
		Where("hotels_amenities.hotel_id = ?", hotelId).
		Find(&amenities)
	log.Debug("Amenities by Hotel ID: ", amenities)

	return amenities
}

func DeleteAmenitieById(amenitieId int) e.ApiError {

	err := Db.Delete(&model.Amenitie{}, amenitieId).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewBadRequestApiError("Amenitie not found")
		}
		return e.NewBadRequestApiError("Failed to delete amenitie")
	}

	return nil
}