package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/sannonthachai/poc-socket-go/socket"

	socketio "github.com/googollee/go-socket.io"
)

func socketMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
		c.Response().Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Response().Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		c.Request().Header.Del("Origin")
		return next(c)
	}
}

func main() {

	server, _ := socketio.NewServer(nil)
	socket := socket.NewSocket(server)

	socket.InitSocket()

	go server.Serve()
	defer server.Close()

	e := echo.New()
	e.HideBanner = true

	e.Static("/", "../asset")

	group := e.Group("/test")
	group.Any("/socket.io/", func(context echo.Context) error {
		context.Request().Header.Del("Origin")
		server.ServeHTTP(context.Response(), context.Request())
		return nil
	})

	e.GET("/broadcast", func(context echo.Context) error {
		socket.Server.BroadcastToRoom("/", "sentToCustodianRoom", "sentToCustodian", "TEST")
		return nil
	})

	e.Logger.Fatal(e.Start("localhost:8080"))

}

func socketConnect(server *socketio.Server) {
	server.OnConnect("", socketSentToCustodianInitial)
}

func socketDisconnect(server *socketio.Server) {
	server.OnDisconnect("", func(s socketio.Conn, reason string) {
		fmt.Println(s.Context())
		fmt.Println("closed", reason)
	})
}

func socketSentToCustodianInitial(s socketio.Conn) error {
	s.SetContext("test")
	fmt.Println("connected:", s.ID())
	s.Join("bcast")
	return nil
}

// e.GET("/connect", func(context echo.Context) error {
// 	socket.SentToCustodianInitSocket()
// 	return nil
// })

// server.OnError("/", func(s socketio.Conn, e error) {
// 	fmt.Println("meet error:", e)
// })

// server.OnDisconnect("/", func(s socketio.Conn, reason string) {
// 	fmt.Println(s.Context())
// 	fmt.Println("closed", reason)
// })

// server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
// 	fmt.Println("notice:", msg)
// 	s.Emit("reply", "have "+msg)
// })

// server.OnEvent("/", "msg", func(s socketio.Conn, msg string) string {
// 	s.SetContext(msg)
// 	return "recv " + msg
// })

// server.OnEvent("/", "echo", func(s socketio.Conn, msg interface{}) {
// 	s.Emit("echo", msg)
// })

// server.OnEvent("/", "bye", func(s socketio.Conn) string {
// 	last := s.Context().(string)
// 	s.Emit("bye", last)
// 	s.Close()
// 	return last
// })
