package api

import (
	"encoding/json"
	"net/http"
	"securechat/backend/src/controller/model"
	"securechat/backend/src/db/schema"
	"securechat/backend/src/handler"
	"securechat/backend/src/middleware"
	"securechat/backend/src/service"
)

func RequestChatSession(w http.ResponseWriter, r *http.Request) {
	var request model.RequestSessionRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		handler.ErrorResponse("Invalid request body", &err, w, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	user := r.Context().Value(middleware.UserContextKey).(*schema.User)

	session, err := service.NewChatSessionService(user, request.Email)
	if err != nil {
		handler.ErrorResponse("Failed to request chat session", &err, w, http.StatusInternalServerError)
		return
	}
	handler.SuccessResponse("Chat session requested successfully", session, w, http.StatusCreated)
}

func CreateChatSession(w http.ResponseWriter, r *http.Request) {
	var request model.CreateSessionRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		handler.ErrorResponse("Invalid request body", &err, w, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	user := r.Context().Value(middleware.UserContextKey).(*schema.User)
	session, err := service.CreateChatSession(*user, request)
	if err != nil {
		handler.ErrorResponse("Failed to create chat session", &err, w, http.StatusInternalServerError)
		return
	}
	handler.SuccessResponse("Chat session created successfully", session, w, http.StatusCreated)
}
