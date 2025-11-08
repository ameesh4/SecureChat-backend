package api

import (
	"encoding/json"
	"net/http"
	"securechat/backend/src/controller/model"
	"securechat/backend/src/db/repository"
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

	second_user, err := repository.GetUserByEmail(request.Email)
	if err != nil {
		handler.ErrorResponse("Failed to get second user", &err, w, http.StatusInternalServerError)
		return
	}

	user2Socket, found := handler.Connections[second_user.Id]

	if found {
		session2 := schema.ChatSession{
			Id:           session.Id,
			CreatedAt:    session.CreatedAt,
			UpdatedAt:    session.UpdatedAt,
			User1:        session.User2,
			User2:        session.User1,
			Participant1: session.Participant2,
			Participant2: session.Participant1,
			A1:           session.A2,
			A2:           session.A1,
		}
		user2Socket.Emit("new_chat_session", session2)
	}

}

func GetAllChatSessions(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(middleware.UserContextKey).(*schema.User)
	sessions, err := service.GetAllChatSessions(user.Id)
	if err != nil {
		handler.ErrorResponse("Failed to get all chat sessions", &err, w, http.StatusInternalServerError)
		return
	}
	handler.SuccessResponse("All chat sessions retrieved successfully", sessions, w, http.StatusOK)
}
