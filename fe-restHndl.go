package main

import (
	"encoding/json"
	"net/http"

	"github.com/garekkream/bookshelf/Settings"
)

func GetShelfs(w http.ResponseWriter, r *http.Request) {
	shelfs := Settings.GetConfig().Shelfs

	b, _ := json.Marshal(shelfs)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(b)
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	type verStruct struct {
		Version string `json:"version"`
	}

	v := verStruct{Version: date + ver}
	b, _ := json.Marshal(v)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(b)
}
