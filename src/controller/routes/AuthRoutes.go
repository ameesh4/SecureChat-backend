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
	}
}
