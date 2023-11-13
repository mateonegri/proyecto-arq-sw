package dto_test

import (
	"mvc-proyecto-arq-sw/dto"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookingDto(t *testing.T) {
	// Crear una instancia del DTO de Booking, si modifico alguna y deja de ser igual, da la alerta
	bookingDto := dto.BookingDto{
		Id:       1,
		StartDate: 20240315,
		EndDate: 20240325,
		UserId:   1,
		HotelId:  1,
	}

	// Verificar los valores de los campos del DTO de Booking
	assert.Equal(t, 1, bookingDto.Id, "El ID de la reserva no coincide")
	assert.Equal(t, 20240315, bookingDto.StartDate, "La fecha de inicio no coincide")
	assert.Equal(t, 20240325, bookingDto.EndDate, "La fecha de fin no coincide")
	assert.Equal(t, 1, bookingDto.UserId, "El ID del usuario no coincide")
	assert.Equal(t, 1, bookingDto.HotelId, "El ID del hotel no coincide")
}