package services

import (
	log "github.com/sirupsen/logrus"
	bookingClient "mvc-proyecto-arq-sw/clients/booking"
	hotelClient "mvc-proyecto-arq-sw/clients/hotel"
	userClient "mvc-proyecto-arq-sw/clients/user"
	"mvc-proyecto-arq-sw/dto"
	"mvc-proyecto-arq-sw/model"
	e "mvc-proyecto-arq-sw/utils/errors"
	"fmt"
)

type bookingService struct{}

type bookingServiceInterface interface {
	GetBookingById(id int) (dto.BookingDetailDto, e.ApiError)
	GetBookings() (dto.BookingsDetailDto, e.ApiError)
	InsertBooking(bookingDto dto.BookingDto) (dto.BookingDto, e.ApiError)
	GetBookingByHotelIdAndDate(request dto.CheckRoomDto, idHotel int) (dto.Availability, e.ApiError)
	GetBookingsByUserId(id int) (dto.BookingsDetailDto, e.ApiError)
	GetBookingByUserId(id int) (dto.BookingDetailDto, e.ApiError)
}

var (
	BookingService bookingServiceInterface
)

func init() {
	BookingService = &bookingService{}
}

func (s *bookingService) GetBookingById(id int) (dto.BookingDetailDto, e.ApiError) {
	var booking model.Booking = bookingClient.GetBookingById(id)
	var bookingDto dto.BookingDetailDto

	if booking.Id == 0 {
		return bookingDto, e.NewBadRequestApiError("Booking not found")
	}

	/*	bookingDto.StartDay = booking.StartDay
		bookingDto.StartMonth = booking.StartMonth
		bookingDto.StartYear = booking.StartYear
		bookingDto.EndDay = booking.EndDay
	*/
	bookingDto.Id = booking.Id
	bookingDto.StartDate = booking.StartDate
	bookingDto.EndDate = booking.EndDate
	bookingDto.UserId = booking.UserId
	bookingDto.Username = booking.User.UserName
	bookingDto.HotelId = booking.HotelId
	bookingDto.HotelName = booking.Hotel.HotelName
	bookingDto.Address = booking.Hotel.Address

	return bookingDto, nil
}

func (s *bookingService) GetBookings() (dto.BookingsDetailDto, e.ApiError) {

	var bookings model.Bookings = bookingClient.GetBookings()
	var bookingsDto dto.BookingsDetailDto

	for _, booking := range bookings {
		var bookingDto dto.BookingDetailDto
		id := booking.Id

		bookingDto, _ = s.GetBookingById(id)

		bookingsDto = append(bookingsDto, bookingDto)
	}

	return bookingsDto, nil
}

func (s *bookingService) InsertBooking(bookingDto dto.BookingDto) (dto.BookingDto, e.ApiError) {
	fmt.Printf("entro al service del insert booking")
	var booking model.Booking
	/*
		booking.StartDay = bookingDto.StartDay
		booking.StartMonth = bookingDto.StartMonth
		booking.StartYear = bookingDto.StartYear
		booking.EndDay = bookingDto.EndDay
		booking.EndMonth = bookingDto.EndMonth
		booking.EndYear = bookingDto.EndYear*/

	if userClient.CheckUserById(bookingDto.UserId) == false {
		return bookingDto, e.NewBadRequestApiError("El usuario no esta registrado en el sistema")
	}

	if hotelClient.CheckHotelById(bookingDto.HotelId) == false {
		return bookingDto, e.NewBadRequestApiError("El hotel no esta registrado en el sistema")
	}

	// En caso de que se quiera hardcodear un booking hay que checkear disponibilidad
	var checkAvailabilityDto dto.CheckRoomDto

	checkAvailabilityDto.StartDate = bookingDto.StartDate
	checkAvailabilityDto.EndDate = bookingDto.EndDate

	var responseAvailabilityDto dto.Availability

	responseAvailabilityDto, _ = s.GetBookingByHotelIdAndDate(checkAvailabilityDto, bookingDto.HotelId)

	if responseAvailabilityDto.OkToBook == false {
		return bookingDto, e.NewBadRequestApiError("El hotel no tiene disponibilidad en esas fechas")
	}

	//stringStartDate, _ := strconv.Itoa(intStartDate)
	//stringEndDate, _ := strconv.Itoa(intEndDate)

	//intStartDate, _ := strconv.Atoi(bookingDto.StartDate)
	//intEndDate, _ := strconv.Atoi(bookingDto.EndDate)

	booking.StartDate = bookingDto.StartDate
	booking.EndDate = bookingDto.EndDate
	booking.UserId = bookingDto.UserId
	booking.HotelId = bookingDto.HotelId

	booking = bookingClient.InsertBooking(booking)

	bookingDto.Id = booking.Id

	return bookingDto, nil
}

func (s *bookingService) GetBookingByHotelIdAndDate(request dto.CheckRoomDto, idHotel int) (dto.Availability, e.ApiError) {
	fmt.Println("ENtro al service service ")
	var ocuppiedRoomsDay int = 0


	startDate := request.StartDate
	endDate := request.EndDate

	var hotel model.Hotel = hotelClient.GetHotelById(idHotel)
	var responseDto dto.Availability

	if hotel.Id == 0 {
		return responseDto, e.NewBadRequestApiError("El hotel no se encuentra en el sistema")
	}

	for i := startDate; i < endDate; i = i + 1 {
		ocuppiedRoomsDay = bookingClient.GetBookingsByHotelIdAndUser(idHotel, i)
		log.Debug("Rooms: ", ocuppiedRoomsDay)
		log.Debug("Date: ", i)

		if ocuppiedRoomsDay >= hotel.Rooms {

			responseDto.OkToBook = false

			return responseDto, nil
		}

	}

	responseDto.OkToBook = true

	return responseDto, nil

}

func (s *bookingService) GetBookingByUserId(id int) (dto.BookingDetailDto, e.ApiError) {
	var booking model.Booking = bookingClient.GetBookingByUserId(id)
	var bookingDto dto.BookingDetailDto

	if booking.Id == 0 {
		return bookingDto, e.NewBadRequestApiError("Booking not found")
	}

	/*	bookingDto.StartDay = booking.StartDay
		bookingDto.StartMonth = booking.StartMonth
		bookingDto.StartYear = booking.StartYear
		bookingDto.EndDay = booking.EndDay
	*/
	bookingDto.Id = booking.Id
	bookingDto.StartDate = booking.StartDate
	bookingDto.EndDate = booking.EndDate
	bookingDto.UserId = booking.UserId
	bookingDto.Username = booking.User.UserName
	bookingDto.HotelId = booking.HotelId
	bookingDto.HotelName = booking.Hotel.HotelName
	bookingDto.Address = booking.Hotel.Address

	return bookingDto, nil
}

func (s *bookingService) GetBookingsByUserId(id int) (dto.BookingsDetailDto, e.ApiError) {

	//var bookings model.Bookings = bookingClient.GetBookingsByUserId(id)
	var bookings model.Bookings = bookingClient.GetBookings()
	var bookingsDto dto.BookingsDetailDto

	for _, booking := range bookings {
		var bookingDto dto.BookingDetailDto

		if booking.UserId == id {

			bookingDto, _ = s.GetBookingById(booking.Id)
			bookingsDto = append(bookingsDto, bookingDto)

		}
	}

	return bookingsDto, nil
}
