package main

import (
	"net/http"

	"github.com/garekkream/bookshelf/Settings"
	"github.com/googollee/go-socket.io"
)

func websocketHandlers(server *socketio.Server) {
	server.On("getVersion", hndlVersion)
}

func websocketInit() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		Settings.Log().Fatalf("Failed to create socketio: %s", err)
	}

	websocketHandlers(server)

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./Web")))
	go func() {
		Settings.Log().Fatal(http.ListenAndServe(":1234", nil))
	}()
}
