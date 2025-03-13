package api

import (
	"encoding/json"
	"net/http"
	"note-taking-app/db"

	"github.com/gorilla/mux"
)

func handleGetProfile(w http.ResponseWriter, r *http.Request) {
	username, err := GetUsernameFromCookie(r)
	if err != nil || username == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := db.GetProfile(username)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func RegisterProfileRoutes(r *mux.Router) {
	userRouter := r.PathPrefix("/api/profile").Subrouter()

	userRouter.Use(AuthCheckerMiddleWare)
	userRouter.HandleFunc("/me", handleGetProfile).Methods("GET")
}
