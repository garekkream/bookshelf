package main

import (
	"github.com/garekkream/bookshelf/Settings"
	"github.com/googollee/go-socket.io"
)

func hndlVersion(so socketio.Socket, err error) {
	so.Emit("setVersion", date+ver)
}

func hndlDebugMode(so socketio.Socket, err error) {
	so.Emit("setDebugMode", Settings.GetDebugMode())
}
