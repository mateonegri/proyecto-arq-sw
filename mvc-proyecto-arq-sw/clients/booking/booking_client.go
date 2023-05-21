package booking

import (
	"mvc-proyecto-arq-sw/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

// GetBookingById sirve para el caso de lista de
// reservas por user
func GetBookingById(id int) model.Booking {
	var booking model.Booking

	Db.Where("id = ?", id).Preload("User").Preload("Hotel").First(&booking)
	log.Debug("Booking: ", booking)

	return booking
}

// Sirve para el caso de la vista del admin
// de todos las reservas con su informacion
func GetBookings() model.Bookings {
	var bookings model.Bookings
	Db.Find(&bookings)

	log.Debug("Bookings: ", bookings)

	return bookings
}

// Cargar una reserva a la DB
func InsertBooking(booking model.Booking) model.Booking {
	result := Db.Create(&booking)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Booking Created: ", booking.Id)
	return booking
}
