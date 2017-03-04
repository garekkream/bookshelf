package main

import (
	"github.com/garekkream/bookshelf/Settings"
)

func hndlVersion(so socketio.Socket, err error) {
	so.Emit("setVersion", date+ver)
}

func hndlDebugMode(so socketio.Socket, err error) {
	so.Emit("setDebugMode", Settings.GetDebugMode())
}

func hndlListShelf(so socketio.Socket, err error) {
	shelfs := Settings.GetConfig().Shelfs
	for _, n := range shelfs {
		var message Settings.ShelfList

		message.Active = n.Active
		message.Name = n.Name
		message.Path = n.Path

		so.Emit("setShelfs", message)
	}
}
