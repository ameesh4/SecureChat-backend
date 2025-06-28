package service

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	Id    uint   `json:"user_id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

//func GenerateToken(userId uint, email string) (string, error) {
//	token := jwt.Token{}
//}
