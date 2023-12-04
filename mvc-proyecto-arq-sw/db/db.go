package db

import (
	amenitieClient "mvc-proyecto-arq-sw/clients/amenities"
	bookingClient "mvc-proyecto-arq-sw/clients/booking"
	hotelClient "mvc-proyecto-arq-sw/clients/hotel"
	userClient "mvc-proyecto-arq-sw/clients/user"
	imageClient "mvc-proyecto-arq-sw/clients/image"
	"mvc-proyecto-arq-sw/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "tohotelproyecto"
	DBUser := "root"
	DBPass := ""
	//DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "mysql"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all Clients that we build
	userClient.Db = db
	hotelClient.Db = db
	bookingClient.Db = db
	amenitieClient.Db = db
	imageClient.Db = db 

}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Booking{})
	db.AutoMigrate(&model.Hotel{})
	db.AutoMigrate(&model.Amenities{})
	db.AutoMigrate(&model.Image{})

	log.Info("Finishing Migration Database Tables")
}
