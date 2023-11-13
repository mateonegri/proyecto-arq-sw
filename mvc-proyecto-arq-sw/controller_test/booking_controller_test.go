package controller_test

import (

	"mvc-proyecto-arq-sw/dto"
	service "mvc-proyecto-arq-sw/services"
	"mvc-proyecto-arq-sw/utils/errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	bookingController "mvc-proyecto-arq-sw/controllers/booking"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type TestBookings struct {
}

func (t *TestBookings) InsertBooking(bookingDto dto.BookingDto) (dto.BookingDto, errors.ApiError) {
	if bookingDto.UserId == 0 {
		return dto.BookingDto{}, errors.NewApiError("Error al insertar la reserva", "booking_insert_error", http.StatusInternalServerError, nil)
	}

	return dto.BookingDto{}, nil
}
func (t *TestBookings) GetBookingByHotelIdAndDate(request dto.CheckRoomDto, idHotel int) (dto.Availability, errors.ApiError) {
	return dto.Availability{}, nil
}


func (t *TestBookings) GetBookingByUserId(id int) (dto.BookingDetailDto, errors.ApiError) {
	if id == 1 {
		return dto.BookingDetailDto{
			Id:       1,
			StartDate: 20240315,
			EndDate: 20210325,
			UserId:   1,
			HotelId:  2,
		}, nil
	}

	return dto.BookingDetailDto{}, errors.NewNotFoundApiError("Booking not found")
}

func (t *TestBookings) GetBookingById(id int) (dto.BookingDetailDto, errors.ApiError) {
	if id == 1 {
		return dto.BookingDetailDto{
			Id:       1,
			StartDate: 20240315, 
			EndDate:   20240325,
			UserId:   1,
			HotelId:  2,
		}, nil
	}

	return dto.BookingDetailDto{}, errors.NewApiError("Booking not found", "booking_not_found", http.StatusNotFound, nil)
}

func (t *TestBookings) GetBookings() (dto.BookingsDetailDto, errors.ApiError) {
	return dto.BookingsDetailDto{}, nil
}


func (t *TestBookings) GetBookingsByUserId(id int) (dto.BookingsDetailDto, errors.ApiError) {
	return dto.BookingsDetailDto{}, nil
}


func TestInsertBooking(t *testing.T) {
	service.BookingService = &TestBookings{}
	router := gin.Default()

	router.POST("/booking", bookingController.InsertBooking)

	// Solicitud HTTP POST - Si se cambia el User id a 0 se ve el error
	myJson := `{
		"hotel_id": 2,
		"date_from": "2024/03/15",
		"date_to": "2024/03/25",
		"user_id": 1
		}`

	response := httptest.NewRecorder()

	var checkAvailabilityDto dto.CheckRoomDto

	checkAvailabilityDto.StartDate = 20240315
	checkAvailabilityDto.EndDate = 20240325

	var responseAvailabilityDto dto.Availability

	if responseAvailabilityDto.OkToBook == false  {
		assert.Equal(t, http.StatusOK, response.Code, "No se pudo lograr la disponibilidad para las fechas especificadas")
	} else {

		bodyJson := strings.NewReader(myJson)
		request, _ := http.NewRequest("POST", "/booking", bodyJson)

		router.ServeHTTP(response, request)

		fmt.Println(response.Body.String())

		// Verificar el código de estado de la respuesta
		assert.Equal(t, http.StatusCreated, response.Code, "El código de respuesta no es el esperado")
	}
}

func TestGetBookingByUserId(t *testing.T) {
	service.BookingService = &TestBookings{}
	router := gin.Default()

	router.GET("/booking/:id", bookingController.GetBookingById)

	// Crear una solicitud HTTP de tipo GET al endpoint /booking/{id}

	// Si se cambia el id a otro numero se ve el error
	request, _ := http.NewRequest("GET", "/booking/1", nil)

	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	fmt.Println(response.Body.String())

	// Verificar el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, response.Code, "El ID buscado no existe")
}