package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Id uint `json:"user_id"`
	jwt.RegisteredClaims
}

type JWTService struct {
	SecretKey []byte
	issuer    string
}

func NewJWTService(secretKey []byte, issuer string) *JWTService {
	return &JWTService{
		SecretKey: secretKey,
		issuer:    issuer,
	}
}

func (J *JWTService) GenerateToken(userId uint) (string, error) {
	claims := CustomClaims{
		Id: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(14 * 24 * time.Hour)), // 14 days
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    J.issuer,
			Subject:   "user_auth",
			ID:        fmt.Sprintf("%d-%d", userId, time.Now().Unix()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(J.SecretKey)
}

func (J *JWTService) GenerateRefreshToken(userId uint) (string, error) {
	claims := CustomClaims{
		Id: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)), // 7 days
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    J.issuer,
			Subject:   "user_auth_refresh",
			ID:        fmt.Sprintf("%d-%d", userId, time.Now().Unix()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(J.SecretKey)
}

func (J *JWTService) ValidateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return J.SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}

func (J *JWTService) ValidateRefreshToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return J.SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid refresh token")
	}
}

func (J *JWTService) ExtractUserIdFromToken(tokenString string) (uint, error) {
	claims, err := J.ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}
	return claims.Id, nil
}

func (J *JWTService) isTokenExpired(tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return J.SecretKey, nil
	})
	if err != nil {
		return false, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims.ExpiresAt.Time.Before(time.Now()), nil
	} else {
		return false, fmt.Errorf("invalid token")
	}
}
