package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	fmt.Println("Hello World")

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		log.Println("New connection")
	})

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	http.Handle("/scoket.io/", server)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
