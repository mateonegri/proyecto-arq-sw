package imageController

import (
	"mvc-proyecto-arq-sw/dto"
	service "mvc-proyecto-arq-sw/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetImageById(c *gin.Context) {
	log.Debug("Image id to load: " + c.Param("id"))

	var imageDto dto.ImageDto
	id, _ := strconv.Atoi(c.Param("id"))
	imageDto, err := service.ImageService.GetImageById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, imageDto)
}

func GetImages(c *gin.Context) {
	var imagesDto dto.ImagesDto
	imagesDto, err := service.ImageService.GetImages()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, imagesDto)
}

func ImageInsert(c *gin.Context) {
	var imageDto dto.ImageDto
	// Obtener el ID del hotel del contexto o de los parámetros
	hotelID, erint := strconv.Atoi(c.Param("id"))
	if erint != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erint.Error()})
		return
	}

	// Obtener los datos de la imagen del cuerpo de la solicitud
	imageFile, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Guardar la imagen y manejar la lógica de relación con el hotel
	imageDto, er := service.ImageService.ImageInsert(hotelID, imageFile)
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, imageDto)
}

func GetImagesByHotelId(c *gin.Context) {
	log.Debug("Hotel id to load images: " + c.Param("id"))

	hotelID, _ := strconv.Atoi(c.Param("id"))
	imagesDto, err := service.ImageService.GetImagesByHotelId(hotelID)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, imagesDto)
}

func DeleteImageById(c *gin.Context) {
	// Obtiene el ID de la imagen de los parámetros de la solicitud
	imageId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
		return
	}

	// Llama al servicio para eliminar la imagen por su ID
	err = service.ImageService.DeleteImageById(imageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}