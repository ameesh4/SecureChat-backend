package routes

import (
	"log"
	"net/http"
	"securechat/backend/src/middleware"
	"strings"
)

type Middleware func(http.Handler) http.Handler

func Chain(middlewares ...Middleware) Middleware {
	return func(final http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}
		return final
	}
}

func Router() *http.ServeMux {
	log.Printf("Initializing router...")
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/auth/", apiRouter)
	mux.Handle("/api/v1/backend/", Chain(middleware.AuthMiddleware)(http.HandlerFunc(apiRouter)))
	mux.Handle("/api/v1/admin/", Chain(middleware.AdminMiddleware)(http.HandlerFunc(apiRouter)))
	mux.Handle("/api/v1/chat-session/", Chain(middleware.AuthMiddleware)(http.HandlerFunc(apiRouter)))
	mux.Handle("/api/v1/profile", Chain(middleware.AuthMiddleware)(http.HandlerFunc(apiRouter)))
	mux.Handle("/api/v1/chat-message", Chain(middleware.AuthMiddleware)(http.HandlerFunc(apiRouter)))

	return mux
}

func apiRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/v1")
	switch {
	case strings.HasPrefix(path, "/auth"):
		AuthRoutes(w, r, strings.TrimPrefix(path, "/auth"))
	case strings.HasPrefix(path, "/admin"):
		AdminRoutes(w, r, strings.TrimPrefix(path, "/admin"))
	case strings.HasPrefix(path, "/chat-session"):
		ChatSessionRoutes(w, r, strings.TrimPrefix(path, "/chat-session"))
	case strings.HasPrefix(path, "/profile"):
		ProfileRoutes(w, r, strings.TrimPrefix(path, "/profile"))
	case strings.HasPrefix(path, "/chat-message"):
		ChatMessageRoutes(w, r, strings.TrimPrefix(path, "/chat-message"))
	default:
		http.NotFound(w, r)
	}
}
