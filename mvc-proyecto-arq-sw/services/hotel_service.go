package services

import (
	hotelClient "mvc-proyecto-arq-sw/clients/hotel"

	"mvc-proyecto-arq-sw/dto"
	"mvc-proyecto-arq-sw/model"
	e "mvc-proyecto-arq-sw/utils/errors"
)

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotels() (dto.HotelsDto, e.ApiError)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
	GetHotelById(id int) (dto.HotelDto, e.ApiError)
	GetAvailableRoomsById(id int) (dto.HotelDto, e.ApiError)
	//UpdateAvailableRooms --> Faltaria esta funcion, aca y en el client.
	//Le paso el dia y en base a eso me consulta la availability
	//CheckAvailabiltyRange()

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

	hotelDto.HotelName = hotel.HotelName
	hotelDto.HotelDescription = hotel.HotelDescription
	hotelDto.Address = hotel.Address
	hotelDto.Rooms = hotel.Rooms
	hotelDto.ImageURL = hotel.ImageURL
	hotelDto.AvailableRooms = hotel.AvailableRooms

	return hotelDto, nil

}

func (s *hotelService) GetHotels() (dto.HotelsDto, e.ApiError) {

	var hotels model.Hotels = hotelClient.GetHotels()
	var hotelsDto dto.HotelsDto

	for _, hotel := range hotels {
		var hotelDto dto.HotelDto

		hotelDto.HotelName = hotel.HotelName
		hotelDto.HotelDescription = hotel.HotelDescription
		hotelDto.Address = hotel.Address
		hotelDto.Rooms = hotel.Rooms
		hotelDto.Id = hotel.Id
		hotelDto.AvailableRooms = hotel.AvailableRooms

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}

func (s *hotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel

	hotel.Id = hotelDto.Id
	hotel.HotelName = hotelDto.HotelName
	hotel.HotelDescription = hotelDto.HotelDescription
	hotel.Address = hotelDto.Address
	hotel.Rooms = hotelDto.Rooms
	hotel.AvailableRooms = hotelDto.AvailableRooms
	hotel.ImageURL = hotelDto.ImageURL

	hotel = hotelClient.InsertHotel(hotel)

	hotelDto.Id = hotel.Id

	return hotelDto, nil
}

func (s *hotelService) GetAvailableRoomsById(id int) (dto.HotelDto, e.ApiError) {
	var hotel model.Hotel = hotelClient.GetHotelById(id)
	var hotelDto dto.HotelDto

	if hotel.Id == 0 {
		return hotelDto, e.NewBadRequestApiError("Hotel no encontrado")
	}

	hotelDto.AvailableRooms = hotel.AvailableRooms

	return hotelDto, nil
}
