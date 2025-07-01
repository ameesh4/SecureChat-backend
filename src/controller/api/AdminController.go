package api

import (
	"net/http"
	"securechat/backend/src/handler"
	"securechat/backend/src/service"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetAllUsers()
	if err != nil {
		handler.ErrorResponse("Failed to retrieve users", w, http.StatusInternalServerError)
		return
	}
	handler.SuccessResponse("Users retrieved successfully", users, w, http.StatusOK)
}
