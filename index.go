package main

import (
	"fmt"
	"log"
	"net/http"
	"securechat/backend/src/controller/routes"
	"securechat/backend/src/db"
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
	handler := routes.Router()
	handler.HandleFunc("/", helloHandler)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
