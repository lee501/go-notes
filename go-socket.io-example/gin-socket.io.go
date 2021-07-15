package main

import (
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func CorsMiddleware(allowUrl string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Accept-COntrol-Allow-Origin", allowUrl)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Request.Header.Del("Origin")

		c.Next()
	}
}

//socket.io demo
func main() {
	router := gin.New()
	//socket server
	server := socketio.NewServer(nil)
	//socket method
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected: ", s.ID())
		return nil
	})
	server.OnEvent("/", "notice", func(s socketio.Conn, msg string) {
		s.Emit("reply", "get "+msg)
	})
	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return msg
	})
	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		log.Println("closed", msg)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatal("socket.io listen err", err)
		}
	}()
	defer server.Close()

	router.Use(CorsMiddleware("localhost:3001"))
	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))

	//http server
	serve := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	if err := serve.ListenAndServe(); err != nil {
		log.Fatal("http listen err", err)
	}
}
