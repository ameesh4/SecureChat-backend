package middleware

import (
	"net/http"
	"os"
	"securechat/backend/src/db/repository"
	"securechat/backend/src/service"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		jwtService := service.NewJWTService([]byte(os.Getenv("JWT_SECRET_KEY")), "securechat")
		_, err := jwtService.ValidateToken(token)
		if err != nil {
			id, err := jwtService.ExtractUserIdFromToken(token)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			user, err := repository.GetUserByID(id)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			_, err = jwtService.ValidateRefreshToken(user.RefreshToken)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			token, err = jwtService.GenerateToken(user.Id)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			w.Header().Set("Authorization", token)
		}
		w.Header().Set("Authorization", token)
		next.ServeHTTP(w, r)
	})
}
