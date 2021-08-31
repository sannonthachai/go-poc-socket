package socket

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

type socket struct {
	Server *socketio.Server
}

func NewSocket(server *socketio.Server) *socket {
	return &socket{
		Server: server,
	}
}

func (s *socket) InitSocket() {
	s.Server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("connected:", s.ID())
		s.Join("sentToCustodianRoom")
		return nil
	})

	s.Server.OnEvent("/", "room", func(s socketio.Conn, room string) {
		s.Join(room)
	})

	s.Server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	s.Server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println(s.Namespace())
		fmt.Printf("%s %s", s.Namespace(), reason)
	})
}

func (s *socket) Error() {
	s.Server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
}
