package dao

import (
	"mvc-proyecto-arq-sw/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"errors"
	e "mvc-proyecto-arq-sw/utils/errors"
)

var Db *gorm.DB

func GetImageById(id int) model.Image {
	var image model.Image

	Db.Where("id = ?", id).Preload("Hotel").First(&image)
	log.Debug("Image: ", image)

	return image
}

func GetImages() model.Images {
	var images model.Images
	Db.Preload("Hotel").Find(&images)

	log.Debug("Images: ", images)

	return images
}

func ImageInsert(image model.Image) model.Image {
	result := Db.Create(&image)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Image Created: ", image.Id)
	return image
}

func GetImagesByHotelId(hotelId int) model.Images {
	var images model.Images

	Db.Where("hotel_id = ?", hotelId).Find(&images)
	log.Debug("Images: ", images)

	return images
}

func DeleteImageById(imageId int) e.ApiError {

	err := Db.Delete(&model.Image{}, imageId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewBadRequestApiError("Image not found")
		}
		return e.NewBadRequestApiError("Failed to delete Image")
	}

	return nil
}