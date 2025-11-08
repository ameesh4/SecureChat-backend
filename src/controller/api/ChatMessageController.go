package api

import (
	"net/http"
	"securechat/backend/src/handler"
	"securechat/backend/src/service"
	"strconv"
)

func GetChatMessages(w http.ResponseWriter, r *http.Request) {
	sessionId := r.URL.Query().Get("session_id")
	if sessionId == "" {
		handler.ErrorResponse("Session id is required", nil, w, http.StatusBadRequest)
		return
	}
	sessionIdUint, err := strconv.ParseUint(sessionId, 10, 64)
	if err != nil {
		handler.ErrorResponse("Invalid session id", &err, w, http.StatusBadRequest)
		return
	}
	messages, err := service.GetChatMessages(uint(sessionIdUint))
	if err != nil {
		handler.ErrorResponse("Failed to get chat messages", &err, w, http.StatusInternalServerError)
		return
	}
	handler.SuccessResponse("Chat messages retrieved successfully", messages, w, http.StatusOK)
}
