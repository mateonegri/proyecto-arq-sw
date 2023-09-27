package app

import (
	//userController "mvc-go/controllers/user"

	log "github.com/sirupsen/logrus"
	amenitieController "mvc-proyecto-arq-sw/controllers/amenitie"
	bookingController "mvc-proyecto-arq-sw/controllers/booking"
	hotelController "mvc-proyecto-arq-sw/controllers/hotel"
	userController "mvc-proyecto-arq-sw/controllers/user"
)

func mapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/user", userController.UserInsert) // Sign In

	// Hotels Mapping
	router.GET("/hotel/:id", hotelController.GetHotelById)
	router.GET("/hotel", hotelController.GetHotels)
	router.POST("/hotel", hotelController.InsertHotel)
	router.POST("/hotel/availability/:id", hotelController.CheckAvailability)
	router.PUT("/hotel/update", hotelController.UpdateHotel)
	router.DELETE("/hotel/delete/:hotel_id/:user_id", hotelController.DeleteHotel)
	router.PUT("/hotel/:id/add-amenitie/:id_amenitie", hotelController.AddHotelAmenitie)
	router.DELETE("/hotel/:id/remove-amenitie/:id_amenitie", hotelController.DeleteHotelAmenitie)

	// Bookings Mapping
	router.GET("/booking/:id", bookingController.GetBookingById)
	router.GET("/booking", bookingController.GetBookings)
	router.POST("/booking", bookingController.InsertBooking)
	router.GET("/booking/user/:user_id", bookingController.GetBookingsByUserId)

	// Amenities Mapping
	router.GET("/amenitie/:id", amenitieController.GetAmenitieById)
	router.GET("/amenities", amenitieController.GetAmenities)
	router.POST("/amenitie", amenitieController.AmenitieInsert)
	router.GET("/amenities/hotel/:id", amenitieController.GetAmenitiesByHotelId)
	router.DELETE("/amenitie/:id", amenitieController.DeleteAmenitieById)

	// Login
	router.POST("/login", userController.Login)

	log.Info("Finishing mappings configurations")
}
