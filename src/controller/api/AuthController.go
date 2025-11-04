package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"securechat/backend/src/controller/model"
	"securechat/backend/src/db/repository"
	"securechat/backend/src/db/schema"
	"securechat/backend/src/handler"
	"securechat/backend/src/service"

	"github.com/joho/godotenv"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user model.RegisterRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		handler.ErrorResponse("Invalid request body", &err, w, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	var userSchema schema.User = schema.User{
		Email:    user.Email,
		Password: user.Password,
	}

	response, err := service.RegisterUser(&userSchema)
	if err != nil {
		handler.ErrorResponse("Failed to register user", &err, w, http.StatusBadRequest)
		return
	}
	handler.SuccessResponse("User registered successfully", response, w, http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user model.LoginRequest
	err := godotenv.Load()
	if err != nil {
		handler.ErrorResponse("Error loading environment variables", &err, w, http.StatusInternalServerError)
		return
	}
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		handler.ErrorResponse("JWT secret key not set", &err, w, http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		handler.ErrorResponse("Invalid request body", &err, w, http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			handler.ErrorResponse("Failed to close request body", &err, w, http.StatusInternalServerError)
			return
		}
	}(r.Body)

	jwtService := service.NewJWTService([]byte(secretKey), "securechat")
	authService := service.NewAuthService(jwtService)

	response, err := authService.AuthenticateUser(user)
	if err != nil {
		handler.ErrorResponse("Failed to login", &err, w, http.StatusUnauthorized)
		return
	}
	handler.SuccessResponse("Login successful", response, w, http.StatusOK)
}

func ValidateToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		handler.ErrorResponse("Unauthorized", nil, w, http.StatusUnauthorized)
		return
	}
	token = token[len("Bearer "):]
	jwtService := service.NewJWTService([]byte(os.Getenv("JWT_SECRET_KEY")), "securechat")
	_, err := jwtService.ValidateToken(token)
	if err != nil {
		id, err := jwtService.ExtractUserIdFromToken(token)
		if err != nil {
			handler.ErrorResponse("Unauthorized", &err, w, http.StatusUnauthorized)
			return
		}
		user, err := repository.GetUserByID(id)
		if err != nil {
			handler.ErrorResponse("Unauthorized", &err, w, http.StatusUnauthorized)
			return
		}
		_, err = jwtService.ValidateRefreshToken(user.RefreshToken)
		if err != nil {
			handler.ErrorResponse("Unauthorized", &err, w, http.StatusUnauthorized)
			return
		}
		token, err = jwtService.GenerateToken(user.Id)
		if err != nil {
			handler.ErrorResponse("Unauthorized", &err, w, http.StatusUnauthorized)
			return
		}
		w.Header().Set("Authorization", token)
		user.Password = ""
		user.RefreshToken = ""
		handler.SuccessResponse("Token validated", user, w, http.StatusOK)
	} else {
		// Token is valid, get user from token and set in context
		id, err := jwtService.ExtractUserIdFromToken(token)
		if err != nil {
			handler.ErrorResponse("Unauthorized", &err, w, http.StatusUnauthorized)
			return
		}
		user, err := repository.GetUserByID(id)
		if err != nil {
			handler.ErrorResponse("Unauthorized", &err, w, http.StatusUnauthorized)
			return
		}
		w.Header().Set("Authorization", token)
		user.Password = ""
		user.RefreshToken = ""
		handler.SuccessResponse("Token validated", user, w, http.StatusOK)
	}
}
