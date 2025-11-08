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

type RegisterRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	PublicKey string `json:"public_key"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}
