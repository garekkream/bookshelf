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
	router.Methods("POST").Path("/shelfs").HandlerFunc(AddShelf)
	router.Methods("DELETE").Path("/shelfs/{id}").HandlerFunc(DelShelf)
	router.Methods("GET").Path("/settings").HandlerFunc(GetSettings)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("Web/")))

	go func() {
		Settings.Log().Fatal(http.ListenAndServe(":1234", router))
	}()
}
