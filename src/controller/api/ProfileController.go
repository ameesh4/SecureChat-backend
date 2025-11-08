package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"securechat/backend/src/controller/model"
	"securechat/backend/src/db/repository"
	"securechat/backend/src/db/schema"
	"securechat/backend/src/handler"
	"securechat/backend/src/middleware"
	"time"
)

type UserResponse struct {
	Id        uint   `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserContextKey).(*schema.User)
	userResponse := &UserResponse{
		Id:        user.Id,
		Email:     user.Email,
		Name:      user.Name,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	handler.SuccessResponse("Profile retrieved successfully", userResponse, w, http.StatusOK)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserContextKey).(*schema.User)
	var request model.UpdateUserRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		handler.ErrorResponse("Invalid request body", &err, w, http.StatusBadRequest)
		return
	}
	user.Name = request.Name
	user.Bio = request.Bio
	user.UpdatedAt = time.Now().Unix()
	fmt.Println(user)
	updatedUser, err := repository.UpdateUser(user)
	if err != nil {
		handler.ErrorResponse("Failed to update profile", &err, w, http.StatusInternalServerError)
		return
	}
	userResponse := &UserResponse{
		Id:        updatedUser.Id,
		Email:     updatedUser.Email,
		Name:      updatedUser.Name,
		Bio:       updatedUser.Bio,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	}
	handler.SuccessResponse("Profile updated successfully", userResponse, w, http.StatusOK)
}
