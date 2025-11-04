package routes

import (
	"net/http"
	"securechat/backend/src/controller/api"
	"strings"
)

func AdminRoutes(w http.ResponseWriter, r *http.Request, path string) {
	switch {
	case strings.HasPrefix(path, "/users"):
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		api.GetAllUsers(w, r)
	default:
		http.NotFound(w, r)
	}
}
