package api

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"securechat/backend/src/controller/model"
	"securechat/backend/src/db/schema"
	"securechat/backend/src/handler"
	"securechat/backend/src/service"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user schema.User

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		handler.ErrorResponse("Invalid request body", w, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	response, err := service.RegisterUser(&user)
	if err != nil {
		handler.ErrorResponse(err.Error(), w, http.StatusBadRequest)
		return
	}
	handler.SuccessResponse("User registered successfully", response, w, http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user model.LoginRequest
	err := godotenv.Load()
	if err != nil {
		handler.ErrorResponse("Error loading environment variables", w, http.StatusInternalServerError)
		return
	}
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		handler.ErrorResponse("JWT secret key not set", w, http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		handler.ErrorResponse("Invalid request body", w, http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			handler.ErrorResponse("Failed to close request body", w, http.StatusInternalServerError)
			return
		}
	}(r.Body)

	jwtService := service.NewJWTService([]byte(secretKey), "securechat")
	authService := service.NewAuthService(jwtService)

	response, err := authService.AuthenticateUser(user)
	if err != nil {
		handler.ErrorResponse(err.Error(), w, http.StatusUnauthorized)
		return
	}
	handler.SuccessResponse("Login successful", response, w, http.StatusOK)
}
