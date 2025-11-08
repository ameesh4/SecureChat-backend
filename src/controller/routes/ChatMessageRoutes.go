package routes

import (
	"net/http"
	"securechat/backend/src/controller/api"
)

func ChatMessageRoutes(w http.ResponseWriter, r *http.Request, path string) {
	switch r.Method {
	case http.MethodGet:
		api.GetChatMessages(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
