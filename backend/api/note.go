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
func handleUpdateNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var newNote CreateNote
	err := json.NewDecoder(r.Body).Decode(&newNote)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Println(id, newNote.Content)
	note, err := db.UpdateNote(id, newNote.Content)
	if err != nil {
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(note)
}
func handleDeleteNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := db.DeleteNote(id)
	if err != nil {
		http.Error(w, "Failed to delete note", http.StatusInternalServerError)
		return
	} else {
		fmt.Println(id, "deleted")
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Note deleted successfully"})

}
func RegisterNoteRoutes(r *mux.Router) {
	noteRouter := r.PathPrefix("/api/note").Subrouter()

	noteRouter.Use(AuthCheckerMiddleWare)
	noteRouter.HandleFunc("/all", handleGetNote).Methods("GET")
	noteRouter.HandleFunc("/create", handleCreateNote).Methods("PUT")
	noteRouter.HandleFunc("/update/{id}", handleUpdateNote).Methods("PATCH")
	noteRouter.HandleFunc("/delete/{id}", handleDeleteNote).Methods("DELETE")
}
