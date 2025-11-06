package handler

import (
	"fmt"

	socketio "github.com/zishang520/socket.io/v2/socket"
)

// SocketServer wraps the socket.io server
type SocketServer struct {
	Server *socketio.Server
}

// Close closes the socket server
func (s *SocketServer) Close() {
	s.Server.Close(nil)
}

// InitializeSocket sets up socket.io and returns a wrapped server instance
func InitializeSocket() *SocketServer {
	server := socketio.NewServer(nil, nil)

	server.On("connection", func(clients ...any) {
		client := clients[0].(*socketio.Socket)
		fmt.Println("✅ Client connected:", client.Id())

		// Handle "message" event
		client.On("message", func(args ...any) {
			if len(args) > 0 {
				msg := args[0].(string)
				fmt.Println("📩 Message received:", msg)
				client.Emit("reply", "Server received: "+msg)
			}
		})

		client.On("welcome", func(args ...any) {
			fmt.Println("👋 Welcome message received:", args)
			client.Emit("welcome", "Welcome to the Secure Chat!")
		})

		// Handle disconnection
		client.On("disconnect", func(...any) {
			fmt.Println("❌ Client disconnected:", client.Id())
		})
	})

	return &SocketServer{
		Server: server,
	}
}
