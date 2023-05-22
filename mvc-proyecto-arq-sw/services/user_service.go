package services

import (
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
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

	user.Name = userDto.Name
	user.LastName = userDto.LastName
	user.UserName = userDto.UserName
	user.Password = userDto.Password //Ver como hasheo la pass
	user.Email = userDto.Email

	user = userClient.InsertUser(user)

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
	if user.Password != loginDto.Password && loginDto.Username == user.UserName {
		return loginResponseDto, e.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": loginDto.Username,
		"pass":     loginDto.Password,
	})
	var jwtKey = []byte("secret_key")
	tokenString, _ := token.SignedString(jwtKey)
	if user.Password != tokenString && loginDto.Username != user.UserName {
		return loginResponseDto, e.NewUnauthorizedApiError("Contraseña incorrecta")
	}

	loginResponseDto.UserId = user.Id
	loginResponseDto.Token = tokenString
	log.Debug(loginResponseDto)
	return loginResponseDto, nil
}
