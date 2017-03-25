package main

import (
	"net/http"

	"github.com/garekkream/bookshelf/Settings"
	"github.com/gorilla/mux"
)

func restInit() {
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/version").HandlerFunc(GetVersion)
	router.Methods("GET").Path("/shelfs").HandlerFunc(GetShelfs)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("Web/")))

	go func() {
		Settings.Log().Fatal(http.ListenAndServe(":1234", router))
	}()
}
