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
		LogPath    string `json:"logPath"`
		DebugMode  bool   `json:"debugMode"`
	}

	v := config{ConfigPath: Settings.GetConfigDir(), LogPath: Settings.GetLogPath(), DebugMode: Settings.GetDebugMode()}
	b, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Write(b)
}

func SetSettings(w http.ResponseWriter, r *http.Request) {
	var settings struct {
		DebugMode string `json:"debugMode"`
	}

	json.NewDecoder(r.Body).Decode(&settings)

	if settings.DebugMode == "true" {
		Settings.DebugModeSave(true)
	} else {
		Settings.DebugModeSave(false)
	}

	b, err := json.Marshal(settings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(201)
	w.Write(b)
}

func GetShelfs(w http.ResponseWriter, r *http.Request) {
	shelfs := Settings.GetConfig().Shelfs

	b, err := json.Marshal(shelfs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Write(b)
}

func AddShelf(w http.ResponseWriter, r *http.Request) {
	var err error

	var body struct {
		Name string
		Path string
	}

	var id struct {
		ID string `json:"id"`
	}

	json.NewDecoder(r.Body).Decode(&body)

	id.ID, err = Shelf.NewShelf(body.Name, body.Path)

	b, err := json.Marshal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)

		return
	}

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
	b, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	w.Write(b)
}
