package main

import (
	"net/http"

	"github.com/garekkream/bookshelf/Settings"
)

func serveHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Web/index.html")
}

func websocketInit() {
	http.HandleFunc("/", serveHome)
	go func() {
		Settings.Log().Fatal(http.ListenAndServe(":1234", nil))
	}()
}
