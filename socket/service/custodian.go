package service

import (
	socketio "github.com/googollee/go-socket.io"
)

func CustodianSocketService(server *socketio.Server) {
	server.OnEvent("/custodian", "room", func(s socketio.Conn, room string) {
		s.Join(room)
	})
}
