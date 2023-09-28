package services

import (
	log "github.com/sirupsen/logrus"
	amenitieClient "mvc-proyecto-arq-sw/clients/amenities"
	hotelClient "mvc-proyecto-arq-sw/clients/hotel"
	userClient "mvc-proyecto-arq-sw/clients/user"

	"mvc-proyecto-arq-sw/dto"
	"mvc-proyecto-arq-sw/model"
	e "mvc-proyecto-arq-sw/utils/errors"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotels() (dto.HotelsDto, e.ApiError)
	InsertHotel(hotelDto dto.HandleHotelDto) (dto.HotelDto, e.ApiError)
	GetHotelById(id int) (dto.HotelDto, e.ApiError)
	UpdateHotel(updateHotelDto dto.HandleHotelDto) (dto.HotelDto, e.ApiError)
	DeleteHotel(idHotel int, idUser int) (dto.DeleteHotelResponseDto, e.ApiError)
	AddHotelAmenitie(hotelId, amenitieId int) e.ApiError
	DeleteHotelAmenitie(hotelId, amenitieId int) e.ApiError
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotelById(id int) (dto.HotelDto, e.ApiError) {
	var hotel model.Hotel = hotelClient.GetHotelById(id)
	var hotelDto dto.HotelDto

	if hotel.Id == 0 {
		return hotelDto, e.NewBadRequestApiError("Hotel no encontrado")
	}

	hotelDto.Id = hotel.Id
	hotelDto.HotelName = hotel.HotelName
	hotelDto.HotelDescription = hotel.HotelDescription
	hotelDto.Address = hotel.Address
	hotelDto.Rooms = hotel.Rooms
	hotelDto.ImageURL = hotel.ImageURL

	amenities := make([]string, 0)

	for _, amenity := range hotel.Amenities {
		amenities = append(amenities, amenity.Name)
	}

	hotelDto.Amenities = amenities

	return hotelDto, nil

}

func (s *hotelService) GetHotels() (dto.HotelsDto, e.ApiError) {

	var hotels model.Hotels = hotelClient.GetHotels()
	var hotelsDto dto.HotelsDto

	for _, hotel := range hotels {
		var hotelDto dto.HotelDto
		id := hotel.Id

		hotelDto, _ = s.GetHotelById(id)

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}

func (s *hotelService) InsertHotel(hotelDto dto.HandleHotelDto) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel
	var user model.User
	var response dto.HotelDto

	user = userClient.GetUserById(hotelDto.UserId)

	if user.Type == false {
		return response, e.NewBadRequestApiError("El usuario no es administrador")
	}

	hotel.HotelName = hotelDto.HotelName
	hotel.HotelDescription = hotelDto.HotelDescription
	hotel.Address = hotelDto.Address
	hotel.Rooms = hotelDto.Rooms
	hotel.ImageURL = hotelDto.ImageURL

	hotel = hotelClient.InsertHotel(hotel)

	if hotel.Id == 0 {
		return response, e.NewBadRequestApiError("Error al insertar hotel")
	}

	hotelDto.Id = hotel.Id

	return response, nil
}

func (s *hotelService) UpdateHotel(updateHotelDto dto.HandleHotelDto) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel
	var savedHotel dto.HotelDto

	savedHotel, _ = s.GetHotelById(updateHotelDto.Id)

	hotel.HotelName = savedHotel.HotelName
	hotel.HotelDescription = savedHotel.HotelDescription
	hotel.Address = savedHotel.Address
	hotel.Rooms = savedHotel.Rooms
	hotel.ImageURL = savedHotel.ImageURL

	if len(updateHotelDto.HotelName) != 0 {
		log.Debug("Nombre del hotel", updateHotelDto)
		hotel.HotelName = updateHotelDto.HotelName
	}

	if len(updateHotelDto.HotelDescription) != 0 {
		hotel.HotelDescription = updateHotelDto.HotelDescription
	}

	if len(updateHotelDto.Address) != 0 {
		hotel.Address = updateHotelDto.Address
	}

	if updateHotelDto.Rooms != 0 {
		hotel.Rooms = updateHotelDto.Rooms
	}

	if len(updateHotelDto.ImageURL) != 0 {
		hotel.ImageURL = updateHotelDto.ImageURL
	}

	hotel.Id = updateHotelDto.Id

	var hotelDto dto.HotelDto
	var user model.User

	user = userClient.GetUserById(updateHotelDto.UserId)

	if user.Type == false {
		return hotelDto, e.NewBadRequestApiError("El usuario no es administrador")
	}

	if hotelClient.CheckHotelById(hotel.Id) == false {
		return hotelDto, e.NewBadRequestApiError("El hotel no esta registrado en el sistema")
	}

	log.Debug("Id model", hotel.Id)
	log.Debug("Id dto", updateHotelDto.Id)

	hotel = hotelClient.UpdateHotelById(hotel)

	if hotel.Id == 0 {
		return hotelDto, e.NewBadRequestApiError("Error al actualizar hotel")
	}

	hotelDto, _ = s.GetHotelById(hotel.Id)

	return hotelDto, nil

}

func (s *hotelService) DeleteHotel(idHotel int, idUser int) (dto.DeleteHotelResponseDto, e.ApiError) {
	var hotel model.Hotel
	var user model.User
	var response dto.DeleteHotelResponseDto

	// user = userClient.GetUserById(deleteHotelDto.UserId)
	user = userClient.GetUserById(idUser)

	// hotel.Id = deleteHotelDto.HotelId
	hotel.Id = idHotel

	if user.Type == false {
		return response, e.NewBadRequestApiError("El usuario no es administrador")
	}

	if hotelClient.CheckHotelById(hotel.Id) == false {
		return response, e.NewBadRequestApiError("El hotel no esta registrado en el sistema")
	}

	var err error

	response.DeleteConfirm, err = hotelClient.DeleteHotel(hotel) // Si DeleteHotel devuelve true --> Se elimino sin problema

	if !response.DeleteConfirm {
		return response, e.NewInternalServerApiError("Error al borrar hotel de la BD", err)
	}

	return response, nil
}

func (s *hotelService) AddHotelAmenitie(hotelId, amenitieId int) e.ApiError {
	// Obtener el hotel por su ID
	hotel := hotelClient.GetHotelById(hotelId)
	if hotel.Id == 0 {
		return e.NewNotFoundApiError("Hotel not found")
	}

	// Obtener la amenidad por su ID
	amenitie := amenitieClient.GetAmenitieById(amenitieId)
	if amenitie.Id == 0 {
		return e.NewNotFoundApiError("Amenitie not found")
	}

	// Verificar si la amenidad ya está asociada al hotel
	for _, amenity := range hotel.Amenities {
		if amenity.Id == amenitieId {
			return e.NewBadRequestApiError("Amenitie already added to the hotel")
		}
	}

	// Asociar la amenidad al hotel
	hotel.Amenities = append(hotel.Amenities, &amenitie)
	hotelClient.UpdateHotel(hotel)

	return nil
}

func (h *hotelService) DeleteHotelAmenitie(hotelId, amenitieId int) e.ApiError {
	// Obtener el hotel por su ID
	hotel := hotelClient.GetHotelById(hotelId)
	if hotel.Id == 0 {
		return e.NewNotFoundApiError("Hotel not found")
	}

	// Obtener la amenitie por su ID
	amenitie := amenitieClient.GetAmenitieById(amenitieId)
	if amenitie.Id == 0 {
		return e.NewNotFoundApiError("Amenitie not found")
	}

	// Eliminar la amenidad al hotel
	// Encuentra el índice del amenitie que deseas eliminar
	var indexToRemove int = -1
	for i, a := range hotel.Amenities {
		if a.Id == amenitie.Id {
			indexToRemove = i
			break
		}
	}

	// Si se encontró el amenitie, elimínalo de la lista
	if indexToRemove != -1 {
		hotel.Amenities = append(hotel.Amenities[:indexToRemove], hotel.Amenities[indexToRemove+1:]...)
	}

	// Actualiza el hotel en la base de datos
	hotelClient.UpdateHotel(hotel)
	hotelClient.DeleteHotelAmenitie(hotelId, amenitieId)

	return nil
}
