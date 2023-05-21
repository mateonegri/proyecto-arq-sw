package dto

type LoginResponseDto struct {
	UserId int    `json:"user_id"`
	Token  string `json:"token"`
}
