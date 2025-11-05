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
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PublicKey   string `json:"public_key"`
}

type UpdateUserRequest struct {
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	WorkspaceId string `json:"workspace_id"`
}
