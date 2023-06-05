package booking

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"mvc-proyecto-arq-sw/dto"
	service "mvc-proyecto-arq-sw/services"
	"net/http"
	"strconv"
)

func GetBookingById(c *gin.Context) {
	log.Debug("Booking id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var bookingDto dto.BookingDetailDto

	bookingDto, err := service.BookingService.GetBookingById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, bookingDto)
}

func GetBookings(c *gin.Context) {
	var bookingsDto dto.BookingsDetailDto
	bookingsDto, err := service.BookingService.GetBookings()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, bookingsDto)
}

func InsertBooking(c *gin.Context) {
	var bookingDto dto.BookingDto
	err := c.BindJSON(&bookingDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	bookingDto, er := service.BookingService.InsertBooking(bookingDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, bookingDto)
}

func GetBookingsByUserId(c *gin.Context) {
	log.Debug("user id to load: " + c.Param("user_id"))

	id, _ := strconv.Atoi(c.Param("user_id"))

	var bookingsDto dto.BookingsDetailDto

	bookingsDto, err := service.BookingService.GetBookingsByUserId(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, bookingsDto)
}
