package model

import "securechat/backend/src/db/schema"

type LoginResponse struct {
	Token string      `json:"token"`
	User  schema.User `json:"user"`
}

type LoginRequest struct {
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
