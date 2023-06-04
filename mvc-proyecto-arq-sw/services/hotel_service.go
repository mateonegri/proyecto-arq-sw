package services

import (
	log "github.com/sirupsen/logrus"
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
