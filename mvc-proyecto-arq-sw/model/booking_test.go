package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBooking(t *testing.T) {
	// Configurar el assert
	assert := assert.New(t)

	// Crear valores de prueba para User
	user := User{
		Id:       1,
		Name:     "Lucrecia",
		LastName: "Sanchez",
		UserName: "lucresanchez",
		Password: "password123",
		Email:    "lucresanchez@gmail.com",
		Type: 		false ,
	}

	// Crear valores de prueba para Hotel
	hotel := Hotel{
		Id:           1,
		HotelName:    "Royal Hotel",
		HotelDescription: "Excelente hotel y ubicacion", 
		Rooms :          4, 
		Address:          "Ny, 9898",
		//ImageURL:        "https://shorturl.at/aeor6",
	}

	// Crear una instancia de Booking con valores de prueba
	booking := Booking{
		Id:       1,
		User: user, 
		Hotel:    hotel,
		StartDate: 20240315, 
		EndDate:   20240325,
	}

	// Verificar que StartDate sea el 15 03 2024 
	expectedDateFrom := 20240315 
	assert.Equal(expectedDateFrom, booking.StartDate, "Se espero que la fecha de inicio sea %v", expectedDateFrom)

	// Verificar que EndDate sea e 25 03 2024 
	expectedDateTo := 20240325
	assert.Equal(expectedDateTo, booking.EndDate, "Se espero que la fecha de fin sea %v", expectedDateTo)

	// Verificar que Id de Booking
	assert.Equal(1, booking.Id, "El ID de la reserva no coincide")

	// Verificar propiedades de User
	assert.Equal("Lucrecia", booking.User.Name, "El nombre no coincide")
	assert.Equal("lucresanchez", booking.User.UserName, "El nombre de usuario no coincide")

	// Verificar propiedades de Hotel
	assert.Equal("Royal Hotel", booking.Hotel.HotelName, "El nombre del hotel coincide")
	assert.Equal(4, booking.Hotel.Rooms, "La cantidad de habitaciones no coincide")
}