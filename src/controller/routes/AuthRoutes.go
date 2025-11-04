package routes

import (
	"net/http"
	"securechat/backend/src/controller/api"
	"strings"
)

func AuthRoutes(w http.ResponseWriter, r *http.Request, path string) {
	switch {
	case strings.HasPrefix(path, "/register"):
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		api.Register(w, r)
	case strings.HasPrefix(path, "/login"):
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		api.Login(w, r)
	case strings.HasPrefix(path, "/verify"):
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		api.ValidateToken(w, r)
	default:
		http.NotFound(w, r)
	}
}
