package user

import (
	"mvc-proyecto-arq-sw/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserByUsername(username string) (model.User, error) {
	var user model.User
	result := Db.Where("user_name = ?", username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func GetUserById(id int) model.User {
	var user model.User

	Db.Where("id = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}

//Checkear si existe un usuario en el sistema

func CheckUserById(id int) bool {
	var user model.User
	var exist bool

	result := Db.Where("id = ?", id).First(&user)

	if result.Error != nil {
		exist = false
		return exist
	}

	exist = true
	return exist
}

func GetUsers() model.Users {
	var users model.Users
	Db.Find(&users)

	log.Debug("Users: ", users)

	return users
}

func InsertUser(user model.User) model.User {
	result := Db.Create(&user)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("User Created: ", user.Id)
	return user
}
