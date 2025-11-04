package middleware

import (
	"context"
	"net/http"
	"os"
	"securechat/backend/src/db/repository"
	"securechat/backend/src/handler"
	"securechat/backend/src/service"
)

type userContextKey string

const UserContextKey userContextKey = "user"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			handler.ErrorResponse("Unauthorized", nil, w, http.StatusUnauthorized)
			return
		}
		token = token[len("Bearer "):]
		jwtService := service.NewJWTService([]byte(os.Getenv("JWT_SECRET_KEY")), "securechat")
		_, err := jwtService.ValidateToken(token)
		if err != nil {
			id, err := jwtService.ExtractUserIdFromToken(token)
			if err != nil {
				handler.ErrorResponse("Unauthorized", &err, w, http.StatusUnauthorized)
				return
			}
			user, err := repository.GetUserByID(id)
			if err != nil {
				handler.ErrorResponse("Unauthorized", &err, w, http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), UserContextKey, user)
			_, err = jwtService.ValidateRefreshToken(user.RefreshToken)
			if err != nil {
				handler.ErrorResponse("Unauthorized", &err, w, http.StatusUnauthorized)
				return
			}
			token, err = jwtService.GenerateToken(user.Id)
			if err != nil {
				handler.ErrorResponse("Unauthorized", &err, w, http.StatusUnauthorized)
				return
			}
			w.Header().Set("Authorization", token)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			// Token is valid, get user from token and set in context
			id, err := jwtService.ExtractUserIdFromToken(token)
			if err != nil {
				handler.ErrorResponse("Unauthorized", &err, w, http.StatusUnauthorized)
				return
			}
			user, err := repository.GetUserByID(id)
			if err != nil {
				handler.ErrorResponse("Unauthorized", &err, w, http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), UserContextKey, user)
			w.Header().Set("Authorization", token)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	})
}
