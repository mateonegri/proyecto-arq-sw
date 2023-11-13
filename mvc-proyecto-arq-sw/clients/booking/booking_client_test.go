package booking

import (
	"testing"

	"mvc-proyecto-arq-sw/model"

	"github.com/stretchr/testify/assert"
)

type MockBookingDAO struct{}

func (m *MockBookingDAO) InsertBooking(booking model.Booking) model.Booking {
	// Simular la lógica de inserción en la base de datos
	// Se establece un ID para la reserva
	booking.Id = 1 // Si modifico a cero, genera la alerta
	return booking
}

// TEST PARA LA FUNCION GETBOOKINGBYIDyy

func TestInsertBooking(t *testing.T) {
	// Crear una instancia del mock del DAO de Booking
	mockDAO := &MockBookingDAO{}

	// Crear una nueva reserva
	newBooking := model.Booking{
		
		UserId:   1,
		HotelId:  1,
		StartDate: 20240315, 
		EndDate:   20240325,
	}

	// Insertar la reserva utilizando el mock del DAO
	inserted := mockDAO.InsertBooking(newBooking)

	// Verificar que la reserva tenga un ID asignado
	assert.NotZero(t, inserted.Id, "La reserva no se pudo realizar")

	// Verificar otros atributos de la reserva
	assert.Equal(t, newBooking.UserId, inserted.UserId)
	assert.Equal(t, newBooking.HotelId, inserted.HotelId)
	assert.Equal(t, newBooking.StartDate, inserted.StartDate)
	assert.Equal(t, newBooking.EndDate, inserted.EndDate)
}

// TEST PARA LA FUNCION GETBOOKINGBYID

func (m *MockBookingDAO) GetBookingByUserId(id int) model.Booking {
	// Simular la búsqueda en la base de datos
	booking := model.Booking{
		Id:       1,
		UserId:   1,
		HotelId:  1,
		StartDate: 20240315, 
		EndDate:   20240325,
	}

	return booking
}

func TestGetBookingByUserId(t *testing.T) {
	// Crear una instancia del mock del DAO de Booking
	mockDAO := &MockBookingDAO{}

	// ID de reserva a buscar - Si la cambio deja de funcionar
	bookingId := 1

	// Obtener la reserva utilizando el mock del DAO
	booking := mockDAO.GetBookingByUserId(bookingId)

	// Verificar que la reserva obtenida tenga el ID correcto
	assert.Equal(t, bookingId, booking.Id, "El ID de la reserva no existe")

	// Verificar otros atributos de la reserva
	assert.Equal(t, 1, booking.UserId)
	assert.Equal(t, 1, booking.HotelId)
	expectedDateFrom := 20240315
	assert.Equal(t, expectedDateFrom, booking.StartDate)
	expectedDateTo := 20240325
	assert.Equal(t, expectedDateTo, booking.EndDate)
}