package main

import (
	"encoding/json"
	"net/http"

	"github.com/garekkream/bookshelf/Settings"
	"github.com/garekkream/bookshelf/Shelf"
	"github.com/gorilla/mux"
)

func GetSettings(w http.ResponseWriter, r *http.Request) {
	type config struct {
		ConfigPath string `json:"configPath"`
		DebugMode  bool   `json:"debugMode"`
	}

	v := config{ConfigPath: Settings.GetConfigDir(), DebugMode: Settings.GetDebugMode()}
	b, _ := json.Marshal(v)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Write(b)
}

func GetShelfs(w http.ResponseWriter, r *http.Request) {
	shelfs := Settings.GetConfig().Shelfs

	b, _ := json.Marshal(shelfs)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Write(b)
}

func AddShelf(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name string
		Path string
	}

	var id struct {
		id string `json:"id"`
	}

	json.NewDecoder(r.Body).Decode(&body)

	id.id = Shelf.NewShelf(body.Name, body.Path)

	b, _ := json.Marshal(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	w.Write(b)
}

func DelShelf(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]

	err := Shelf.DelShelfByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)

		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func GetVersion(w http.ResponseWriter, r *http.Request) {
	type verStruct struct {
		Version string `json:"version"`
	}

	v := verStruct{Version: date + ver}
	b, _ := json.Marshal(v)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Write(b)
}
