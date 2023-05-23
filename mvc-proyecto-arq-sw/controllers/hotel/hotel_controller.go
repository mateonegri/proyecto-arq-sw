package hotelController

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"mvc-proyecto-arq-sw/dto"
	service "mvc-proyecto-arq-sw/services"
	"net/http"
	"strconv"
)

func GetHotelById(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var hotelDto dto.HotelDto

	hotelDto, err := service.HotelService.GetHotelById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}

func GetHotels(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	hotelsDto, err := service.HotelService.GetHotels()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotelsDto)
}

func InsertHotel(c *gin.Context) {
	var hotelDto dto.HotelDto
	err := c.BindJSON(&hotelDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotelDto, er := service.HotelService.InsertHotel(hotelDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}

/* func UpdateHotel(c *gin.Context)  { //Ver si hace falta otro Dto para hacer el update
	var hotelDto dto.HotelDto          //Ya sea de los available rooms, de imagenes, descripciones, etc.

	err := c.BindJSON(&hotelDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}


} */

func CheckAvailability(c *gin.Context) {
	var checkRoomDto dto.CheckRoomDto
	er := c.BindJSON(&checkRoomDto)

	if er != nil {
		log.Error(er.Error())
		c.JSON(http.StatusBadRequest, er.Error())
		return
	}
	log.Debug(checkRoomDto)

	var availabilityResponseDto dto.Availability
	availabilityResponseDto, err := service.BookingService.GetBookingByHotelIdAndDate(checkRoomDto)
	if err != nil {
		if err.Status() == 400 {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusForbidden, err.Error())
		return
	}

	c.JSON(http.StatusOK, availabilityResponseDto)

	//Validar fechas primero. O sino validar en el front
	// Llega el Hotel Id, LLega fechas
	// Busco los Bookings por HotelId y por fecha
	// Y comparo la cantidad de bookings que tengo en esa fecha
	// con la cantidad de habitaciones totales del hotel
	// si en alguna fecha no cumple, no tengo disponibilidad
}
