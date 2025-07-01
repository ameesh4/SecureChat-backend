package middleware

import (
	"context"
	"net/http"
	"os"
	"securechat/backend/src/db/repository"
	"securechat/backend/src/handler"
	"securechat/backend/src/service"
)

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		token = token[len("Bearer "):]
		jwtService := service.NewJWTService([]byte(os.Getenv("JWT_SECRET_KEY")), "securechat")
		_, err := jwtService.ValidateToken(token)
		if err != nil {
			id, err := jwtService.ExtractUserIdFromToken(token)
			if err != nil {
				handler.ErrorResponse("Unauthorized", w, http.StatusUnauthorized)
				return
			}
			user, err := repository.GetUserByID(id)
			if err != nil {
				handler.ErrorResponse("Unauthorized", w, http.StatusUnauthorized)
				return
			}
			if !user.IsAdmin {
				handler.ErrorResponse("Admin Only", w, http.StatusForbidden)
				return
			}
			ctx = context.WithValue(r.Context(), "user", user)
			_, err = jwtService.ValidateRefreshToken(user.RefreshToken)
			if err != nil {
				handler.ErrorResponse("Unauthorized", w, http.StatusUnauthorized)
				return
			}
			token, err = jwtService.GenerateToken(user.Id)
			if err != nil {
				handler.ErrorResponse("Unauthorized", w, http.StatusUnauthorized)
				return
			}
			w.Header().Set("Authorization", token)
		}
		w.Header().Set("Authorization", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
