package model

import "securechat/backend/src/db/schema"

type LoginResponse struct {
	Token string      `json:"token"`
	User  schema.User `json:"user"`
}
