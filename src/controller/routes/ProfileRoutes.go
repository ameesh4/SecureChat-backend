package routes

import (
	"fmt"
	"net/http"
	"securechat/backend/src/controller/api"
)

func ProfileRoutes(w http.ResponseWriter, r *http.Request, path string) {
	fmt.Println("ProfileRoutes", r.Method)
	switch r.Method {
	case http.MethodGet:
		api.GetProfile(w, r)
	case http.MethodPut:
		api.UpdateProfile(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
