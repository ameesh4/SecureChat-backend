package routes

import (
	"net/http"
	"securechat/backend/src/controller/api"
	"strings"
)

func ChatSessionRoutes(w http.ResponseWriter, r *http.Request, path string) {
	switch {
	case strings.HasPrefix(path, "/request"):
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		api.RequestChatSession(w, r)
	case strings.HasPrefix(path, "/create"):
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		api.CreateChatSession(w, r)
	default:
		http.NotFound(w, r)
	}
}
