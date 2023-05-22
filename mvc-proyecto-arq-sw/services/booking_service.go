package services

import (
	bookingClient "mvc-proyecto-arq-sw/clients/booking"
	"mvc-proyecto-arq-sw/dto"
	"mvc-proyecto-arq-sw/model"
	e "mvc-proyecto-arq-sw/utils/errors"
)

type bookingService struct{}

type bookingServiceInterface interface {
	GetBookingById(id int) (dto.BookingDetailDto, e.ApiError)
	GetBookings() (dto.BookingsDetailDto, e.ApiError)
	InsertBooking(bookingDto dto.BookingDto) (dto.BookingDto, e.ApiError)
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

	bookingDto.StartDay = booking.StartDay
	bookingDto.StartMonth = booking.StartMonth
	bookingDto.StartYear = booking.StartYear
	bookingDto.EndDay = booking.EndDay
	bookingDto.EndMonth = booking.EndMonth
	bookingDto.EndYear = booking.EndYear
	bookingDto.UserId = booking.UserId
	bookingDto.Username = booking.User.Name
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

		bookingDto.StartDay = booking.StartDay
		bookingDto.StartMonth = booking.StartMonth
		bookingDto.StartYear = booking.StartYear
		bookingDto.EndDay = booking.EndDay
		bookingDto.EndMonth = booking.EndMonth
		bookingDto.EndYear = booking.EndYear
		bookingDto.UserId = booking.UserId
		bookingDto.Username = booking.User.Name
		bookingDto.HotelId = booking.HotelId
		bookingDto.HotelName = booking.Hotel.HotelName
		bookingDto.Address = booking.Hotel.Address

		bookingsDto = append(bookingsDto, bookingDto)
	}

	return bookingsDto, nil
}

func (s *bookingService) InsertBooking(bookingDto dto.BookingDto) (dto.BookingDto, e.ApiError) {

	var booking model.Booking

	booking.StartDay = bookingDto.StartDay
	booking.StartMonth = bookingDto.StartMonth
	booking.StartYear = bookingDto.StartYear
	booking.EndDay = bookingDto.EndDay
	booking.EndMonth = bookingDto.EndMonth
	booking.EndYear = bookingDto.EndYear
	booking.UserId = bookingDto.UserId
	booking.HotelId = bookingDto.HotelId

	booking = bookingClient.InsertBooking(booking)

	bookingDto.Id = booking.Id

	return bookingDto, nil
}
