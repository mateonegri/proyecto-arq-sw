package services

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	userClient "mvc-proyecto-arq-sw/clients/user"

	"mvc-proyecto-arq-sw/dto"
	"mvc-proyecto-arq-sw/model"
	e "mvc-proyecto-arq-sw/utils/errors"
)

type userService struct{}

type userServiceInterface interface {
	GetUsers() (dto.UsersDto, e.ApiError)
	InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
	GetUserById(id int) (dto.UserDto, e.ApiError)
	Login(loginDto dto.LoginDto) (dto.LoginResponseDto, e.ApiError)
	//GetUserByUsername ?? para el LogIn
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

func (s *userService) GetUserById(id int) (dto.UserDto, e.ApiError) {
	var user model.User = userClient.GetUserById(id)
	var userDto dto.UserDto

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("Usuario no encontrado")
	}

	userDto.Name = user.Name
	userDto.LastName = user.LastName
	userDto.UserName = user.Name
	userDto.Email = user.Email
	userDto.Id = user.Id
	userDto.Type = user.Type

	return userDto, nil

}

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userClient.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto

		//Si el user es Admin no lo incluye en el return

		if !userDto.Type {
			userDto.Name = user.Name
			userDto.LastName = user.LastName
			userDto.UserName = user.UserName
			userDto.Email = user.Email
			userDto.Id = user.Id
			userDto.Type = user.Type
		}

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	var user model.User

	if !userClient.GetUserByEmail(userDto.Email) {
		return userDto, e.NewBadRequestApiError("El email ya esta registrado")
	}

	user.Name = userDto.Name
	user.LastName = userDto.LastName
	user.UserName = userDto.UserName

	var hashedPassword, err = s.HashPassword(userDto.Password)

	if err != nil {
		return userDto, e.NewBadRequestApiError("No se puede utilizar esa contraseña")
	}

	user.Password = hashedPassword //Ver como hasheo la pass
	user.Email = userDto.Email
	user.Type = userDto.Type

	user = userClient.InsertUser(user)

	if user.Id == 0 {
		return userDto, e.NewBadRequestApiError("Nombre de usuario repetido")
	}

	//Validacion de error por repeticion de mail falta

	userDto.Id = user.Id

	return userDto, nil
}

func (s *userService) Login(loginDto dto.LoginDto) (dto.LoginResponseDto, e.ApiError) {

	var user model.User
	user, err := userClient.GetUserByUsername(loginDto.Username)
	var loginResponseDto dto.LoginResponseDto
	loginResponseDto.UserId = -1

	if err != nil {
		return loginResponseDto, e.NewBadRequestApiError("Usuario no encontrado")
	}

	var comparison error = s.VerifyPassword(user.Password, loginDto.Password)

	if loginDto.Username == user.UserName {
		if comparison != nil {
			return loginResponseDto, e.NewUnauthorizedApiError("Contraseña incorrecta 2")
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": loginDto.Username,
		"pass":     loginDto.Password,
	})
	var jwtKey = []byte("secret_key")
	tokenString, _ := token.SignedString(jwtKey)

	var verifyToken error = s.VerifyPassword(user.Password, tokenString)

	if loginDto.Username != user.UserName {
		if verifyToken != nil {
			return loginResponseDto, e.NewUnauthorizedApiError("Contraseña incorrecta 3")
		}
	}

	loginResponseDto.UserId = user.Id
	loginResponseDto.Token = tokenString
	log.Debug(loginResponseDto)
	return loginResponseDto, nil
}

func (s *userService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("No se pudo hashear la password %w", err)
	}

	return string(hashedPassword), nil
}

func (s *userService) VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}
