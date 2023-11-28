package service_test

import (
	"mvc-proyecto-arq-sw/dto"
	// "mvc-proyecto-arq-sw/model"
	service "mvc-proyecto-arq-sw/services"
	e "mvc-proyecto-arq-sw/utils/errors"
	"net/http"
	"testing"
	// "time"
	"fmt"
	"github.com/stretchr/testify/assert"
)

type TestBookings struct {
}

func (t *TestBookings) InsertBooking(bookingDto dto.BookingDto) (dto.BookingDto, e.ApiError) {
	if bookingDto.UserId == 0 {
		return dto.BookingDto{}, e.NewApiError("Error al insertar la reserva", "booking_insert_error", http.StatusInternalServerError, nil)
	}

	return bookingDto, nil
}

func (t *TestBookings) GetBookingByUserId(id int) (dto.BookingDetailDto, e.ApiError) {
	if id == 1 {
		return dto.BookingDetailDto{
			Id:       1,
			StartDate: 20240315,
			EndDate: 20240325,
			UserId:   1,
			HotelId:  1,
		}, nil
	}

	return dto.BookingDetailDto{}, e.NewNotFoundApiError("Booking not found")
}

func (t *TestBookings) GetBookings() (dto.BookingsDetailDto, e.ApiError) {
	return dto.BookingsDetailDto{}, nil
}


func (t *TestBookings) GetBookingsByUserId(id int) (dto.BookingsDetailDto, e.ApiError) {
	return dto.BookingsDetailDto{}, nil
}


func (t *TestBookings) GetBookingByHotelIdAndDate(request dto.CheckRoomDto, idHotel int) (dto.Availability, e.ApiError) {
	return dto.Availability{}, nil
}

func (t *TestBookings) GetBookingById(id int) (dto.BookingDetailDto, e.ApiError) {
	return dto.BookingDetailDto{}, nil
}



func TestInsertBooking(t *testing.T) {
	// Si cambio el valor de los id puedo ver los errores
	booking := dto.BookingDto{
		UserId:   1,
		HotelId:  1,
		StartDate: 20240315, 
		EndDate:   20240325,
	}

	service.BookingService = &TestBookings{}

	
	var checkAvailabilityDto dto.CheckRoomDto

	checkAvailabilityDto.StartDate = booking.StartDate
	checkAvailabilityDto.EndDate = booking.EndDate

	var responseAvailabilityDto dto.Availability

	responseAvailabilityDto, _ = service.BookingService.GetBookingByHotelIdAndDate(checkAvailabilityDto, booking.HotelId)
	fmt.Println("Paso la ida al service")
	if responseAvailabilityDto.OkToBook == false {
		assert.Fail(t, "No se puede realizar la reserva debido a la falta de disponibilidad")
	} else {
		createdBooking, err := service.BookingService.InsertBooking(booking)

		assert.Nil(t, err, "Error al insertar la reserva")
		assert.Equal(t, 1, createdBooking.UserId, "El ID de usuario no coincide")
		assert.Equal(t, 1, createdBooking.HotelId, "El ID de hotel no coincide")
		assert.Equal(t, booking.StartDate, createdBooking.StartDate, "La fecha de inicio no coincide")
		assert.Equal(t, booking.EndDate, createdBooking.EndDate, "La fecha de fin no coincide")
	}



	

}

func TestGetBookingByUserId(t *testing.T) {
	service.BookingService = &TestBookings{}

	// Si cambio los valores de aca puedo ver los errores
	expectedBooking := dto.BookingDto{
		Id:       1,
		UserId:   1,
		HotelId:  1,
		StartDate: 20240315, 
		EndDate:   20240325,
	}

	catchedBooking, err := service.BookingService.GetBookingByUserId(expectedBooking.Id)

	assert.Nil(t, err)
	assert.Equal(t, expectedBooking.Id, catchedBooking.Id, "El ID de reserva no coincide")
	assert.Equal(t, expectedBooking.StartDate, catchedBooking.StartDate, "La fecha de inicio no coincide")
	assert.Equal(t, expectedBooking.EndDate, catchedBooking.EndDate, "La fecha de fin no coincide")
	assert.Equal(t, expectedBooking.UserId, catchedBooking.UserId, "El ID de usuario no coincide")
	assert.Equal(t, expectedBooking.HotelId, catchedBooking.HotelId, "El ID de hotel no coincide")
}