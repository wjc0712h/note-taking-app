package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"note-taking-app/db"

	"github.com/gorilla/mux"
)

type CreateNote struct {
	Content string `json:"content"`
}

func handleCreateNote(w http.ResponseWriter, r *http.Request) {
	username, _ := GetUsernameFromCookie(r)

	var createNote CreateNote
	err := json.NewDecoder(r.Body).Decode(&createNote)

	if err != nil {
		fmt.Println(err)
		return
	}

	note, err := db.CreateNote(createNote.Content, username)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(note)
}

func handleGetNote(w http.ResponseWriter, r *http.Request) {
	username, _ := GetUsernameFromCookie(r)

	notes, err := db.GetNotesbyUsername(username)

	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(notes)
}
func RegisterNoteRoutes(r *mux.Router) {
	noteRouter := r.PathPrefix("/api/note").Subrouter()

	noteRouter.Use(AuthCheckerMiddleWare)
	noteRouter.HandleFunc("/all", handleGetNote).Methods("GET")
	noteRouter.HandleFunc("/", handleCreateNote).Methods("PUT")
}
