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
	var insertHotelDto dto.HandleHotelDto
	err := c.BindJSON(&insertHotelDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var hotelDto dto.HotelDto

	hotelDto, er := service.HotelService.InsertHotel(insertHotelDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}

func UpdateHotel(c *gin.Context) { //Ver si hace falta otro Dto para hacer el update
	var updateHotelDto dto.HandleHotelDto
	err := c.BindJSON(&updateHotelDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var hotelDto dto.HotelDto

	hotelDto, er := service.HotelService.UpdateHotel(updateHotelDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}

func CheckAvailability(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))

	var checkRoomDto dto.CheckRoomDto
	er := c.BindJSON(&checkRoomDto)

	if er != nil {
		log.Error(er.Error())
		c.JSON(http.StatusBadRequest, er.Error())
		return
	}
	log.Debug(checkRoomDto)

	//Validar fechas

	var availabilityResponseDto dto.Availability
	availabilityResponseDto, err := service.BookingService.GetBookingByHotelIdAndDate(checkRoomDto, id)
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

func DeleteHotel(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("hotel_id"))

	id, _ := strconv.Atoi(c.Param("hotel_id"))

	log.Debug("User id to load: " + c.Param("user_id"))

	user_id, _ := strconv.Atoi(c.Param("user_id"))

	/*var deleteHotelDto dto.DeleteHotelDto
	er := c.BindJSON(&deleteHotelDto)

	if er != nil {
		log.Error(er.Error())
		c.JSON(http.StatusBadRequest, er.Error())
		return
	}
	log.Debug(deleteHotelDto)*/

	var responseDeleteDto dto.DeleteHotelResponseDto

	responseDeleteDto, err := service.HotelService.DeleteHotel(id, user_id)

	if err != nil {
		if err.Status() == 400 {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusForbidden, err.Error())
		return
	}

	c.JSON(http.StatusOK, responseDeleteDto)

}
