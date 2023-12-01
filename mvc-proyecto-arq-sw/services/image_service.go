package services

import (
	// "github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	// "path/filepath"

	imageDAO "mvc-proyecto-arq-sw/clients/image"
	"mvc-proyecto-arq-sw/dto"
	"mvc-proyecto-arq-sw/model"

	e "mvc-proyecto-arq-sw/utils/errors"
)

type imageService struct{}

type imageServiceInterface interface {
	GetImageById(id int) (dto.ImageDto, e.ApiError)
	GetImages() (dto.ImagesDto, e.ApiError)
	// ImageInsert(hotelID int, imageFile *multipart.FileHeader) (dto.ImageDto, e.ApiError)
	ImageInsert(hotelID int, URL string) (dto.ImageDto, e.ApiError)
	GetImagesByHotelId(hotelID int) (dto.ImagesDto, e.ApiError)
	DeleteImageById(id int) e.ApiError
}

var (
	ImageService imageServiceInterface
)

func init() {
	ImageService = &imageService{}
}

func (i *imageService) GetImageById(id int) (dto.ImageDto, e.ApiError) {
	image := imageDAO.GetImageById(id)

	if image.Id == 0 {
		return dto.ImageDto{}, e.NewNotFoundApiError("Image not found")
	}

	imageDto := dto.ImageDto{
		Id:  image.Id,
		Url: image.Url,
	}

	return imageDto, nil
}

func (i *imageService) GetImages() (dto.ImagesDto, e.ApiError) {
	images := imageDAO.GetImages()
	imageDtos := make([]dto.ImageDto, len(images))

	for i, image := range images {
		imageDto := dto.ImageDto{
			Id:      image.Id,
			Url:     image.Url,
			HotelId: image.HotelId,
		}
		imageDtos[i] = imageDto
	}

	return dto.ImagesDto{
		Images: imageDtos,
	}, nil
}

/* func (i *imageService) ImageInsert(hotelID int, imageFile *multipart.FileHeader) (dto.ImageDto, e.ApiError) {
	// Crear imageDto para el retorno
	var imageDto dto.ImageDto

	// Generar un nombre único para el archivo de imagen
	fileName := uuid.New().String()

	// Obtener la extensión del archivo
	fileExt := filepath.Ext(imageFile.Filename)

	// Construir la ruta completa del archivo
	filePath := "imagenesHoteles" + "/" + fileName + fileExt

	// Crear una nueva instancia de model.Image
	image := model.Image{
		Url:     filePath,
		HotelId: hotelID,
	}

	// Llamar al DAO de imágenes para insertar la imagen
	image = imageDAO.ImageInsert(image)

	// Guardar el archivo en el directorio correspondiente
	err := saveFile(imageFile, filePath)
	if err != nil {
		// Manejar el error en caso de fallo al guardar la imagen
		_ = i.DeleteImageById(image.Id) // Eliminar la imagen insertada anteriormente
		return imageDto, e.NewInternalServerApiError("Failed to save image", err)
	}

	// Actualizar imageDto con el ID generado
	imageDto.Id = image.Id
	imageDto.Url = image.Url
	imageDto.HotelId = image.HotelId

	return imageDto, nil
} */

func (i *imageService) ImageInsert(hotelID int, URL string) (dto.ImageDto, e.ApiError) {

	var imageDto dto.ImageDto
	var image model.Image

	image.HotelId = hotelID
	image.Url = URL


	image = imageDAO.ImageInsert(image)

	if image.Id == 0 {
		return imageDto, e.NewBadRequestApiError("Error al insertar imagen")
	}

	imageDto.HotelId = image.HotelId
	imageDto.Id = image.Id
	imageDto.Url = image.Url

	return imageDto, nil

}

func saveFile(imageFile *multipart.FileHeader, filePath string) error {
	// Abrir el archivo cargado
	file, err := imageFile.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	// Crear el archivo destino en el sistema de archivos
	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copiar el contenido del archivo cargado al archivo destino
	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}

	return nil
}

func (i *imageService) GetImagesByHotelId(hotelID int) (dto.ImagesDto, e.ApiError) {
	images := imageDAO.GetImagesByHotelId(hotelID)
	imageDtos := make([]dto.ImageDto, len(images))

	for i, image := range images {
		imageDto := dto.ImageDto{
			Id:   image.Id,
			Url:  image.Url,
		}
		imageDtos[i] = imageDto
	}

	return dto.ImagesDto{
		Images: imageDtos,
	}, nil
}

func (i *imageService) DeleteImageById(id int) e.ApiError {
	// Verificar si la imagen existe
	_, err := i.GetImageById(id)
	if err != nil {
		return err
	}

	// Lógica para eliminar la imagen por su ID
	err = imageDAO.DeleteImageById(id)
	if err != nil {
		return e.NewInternalServerApiError("Failed to delete image", err)
	}

	return nil // Sin errores, se eliminó la imagen correctamente
}