package main

import (
	"net/http"
	"note-taking-app/backend/api"
	"note-taking-app/db"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDBConn()
	r := mux.NewRouter().StrictSlash(true)

	api.RegisterNoteRoutes(r)
	api.RegisterProfileRoutes(r)
	api.RegisterAuthRoutes(r)

	http.ListenAndServe(":8080", r)

}
