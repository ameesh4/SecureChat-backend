package api

import (
	"encoding/json"
	"net/http"
	"securechat/backend/src/controller/model"
	"securechat/backend/src/handler"
	"securechat/backend/src/service"
)

func GetChatMessages(w http.ResponseWriter, r *http.Request) {
	var request model.GetChatMessagesRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		handler.ErrorResponse("Invalid request body", &err, w, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	messages, err := service.GetChatMessages(request.SessionId)
	if err != nil {
		handler.ErrorResponse("Failed to get chat messages", &err, w, http.StatusInternalServerError)
		return
	}
	handler.SuccessResponse("Chat messages retrieved successfully", messages, w, http.StatusOK)
}
