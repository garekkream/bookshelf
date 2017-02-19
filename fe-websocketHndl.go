package main

import "github.com/googollee/go-socket.io"

func hndlVersion(so socketio.Socket, err error) {
	so.Emit("setVersion", date+ver)
}
