package main

import (
	"net/http"

	"github.com/garekkream/bookshelf/Settings"
	"github.com/gorilla/mux"
)

func restInit() {
	router := mux.NewRouter().StrictSlash(true)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("Web/")))
	router.Methods("GET").Path("/version").HandlerFunc(GetVersion)

	go func() {
		Settings.Log().Fatal(http.ListenAndServe(":1234", router))
	}()
}
