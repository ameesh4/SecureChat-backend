package handler

import (
	"encoding/json"
	"fmt"
	"securechat/backend/src/controller/model"
	"securechat/backend/src/service"
	"securechat/backend/src/utils"

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

var Connections = make(map[uint]*socketio.Socket)

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

		client.On("auth", func(args ...any) {
			if len(args) > 0 {
				var authRequest model.AuthRequest
				msg := args[0].(string)
				// msg = strings.TrimSpace(msg)
				err := json.Unmarshal([]byte(msg), &authRequest)
				if err != nil {
					fmt.Println("❌ Error decoding auth request:", err)
					return
				}
				fmt.Println("✅ Auth request received:", authRequest.Token)
				user, err := service.ValidateTokenSocket(authRequest.Token)
				if err != nil {
					client.Emit("auth_error", err.Error())
					client.Disconnect(true)
					return
				}
				Connections[user.Id] = client
				client.Emit("auth_success", user)
			}
		})

		client.On("send_message", func(args ...any) {
			if len(args) > 0 {
				var message model.Message
				msg := args[0].(string)
				err := json.Unmarshal([]byte(msg), &message)
				if err != nil {
					fmt.Println("❌ Error decoding message:", err)
					return
				}
				senderId, err := utils.FindKeysByValueConnections(Connections, client)
				if err != nil {
					fmt.Println("❌ Error finding sender ID:", err)
					return
				}
				message.SenderId = senderId[0]
				savedMessage, err := service.SendMessage(message)
				if err != nil {
					client.Emit("message_error", err.Error())
					return
				}
				value, found := Connections[message.ReceiverId]
				if found {
					value.Emit("new_message", savedMessage)
				}
				client.Emit("message_sent", savedMessage)
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
