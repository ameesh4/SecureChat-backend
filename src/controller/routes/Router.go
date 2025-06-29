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

	return mux
}

func apiRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/v1")
	switch {
	case strings.HasPrefix(path, "/auth"):
		AuthRoutes(w, r, strings.TrimPrefix(path, "/auth"))
	}
}
