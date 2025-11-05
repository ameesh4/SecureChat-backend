package main

import (
	"fmt"
	"log"
	"net/http"
	"securechat/backend/src/controller/routes"
	"securechat/backend/src/db"
	"securechat/backend/src/handler"
	"securechat/backend/src/middleware"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello, Secure Chat!")
}

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	socketServer := handler.InitializeSocket()
	defer socketServer.Close()

	mainMux := http.NewServeMux()

	// Register socket.io handler with CORS middleware for initial HTTP handshake
	mainMux.Handle("/socket.io/", middleware.CorsMiddleware(socketServer.Server))

	// Register API routes with CORS middleware
	apiRouter := routes.Router()
	apiRouter.HandleFunc("/", helloHandler)
	mainMux.Handle("/", middleware.CorsMiddleware(apiRouter))

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mainMux))
}
