package handler

import (
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

type SocketServer struct {
	*socketio.Server
}

func InitializeSocket() *SocketServer {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&websocket.Transport{
				CheckOrigin: func(r *http.Request) bool { return true },
			},
		},
	})

	// Default namespace
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		s.Emit("welcome", "Welcome to Secure Chat!")
		return nil
	})

	server.OnEvent("/", "message", func(s socketio.Conn, msg string) {
		s.Emit("message", "Server received: "+msg)
	})

	return &SocketServer{Server: server}
}

func (s *SocketServer) Close() {
	s.Server.Close()
}
