package app

import (
	//userController "mvc-go/controllers/user"

	log "github.com/sirupsen/logrus"
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

	// Bookings Mapping
	router.GET("/booking/:id", bookingController.GetBookingById)
	router.GET("/booking", bookingController.GetBookings)
	router.POST("/booking", bookingController.InsertBooking)

	// Login
	router.POST("/login", userController.Login)

	log.Info("Finishing mappings configurations")
}
