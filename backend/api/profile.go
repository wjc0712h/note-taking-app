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
		http.Error(w, "handleGetProfile(): UNAUTHORIZED ", http.StatusUnauthorized)
		return
	}

	user, err := db.GetProfile(username)
	if err != nil {
		http.Error(w, "handleGetProfile(): NOT FOUND", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func handleCreateProfile(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Username == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	_, err := db.GetProfile(req.Username)
	if err == nil {
		http.Error(w, "handleCreateProfile(): EXIST", http.StatusConflict)
		return
	}

	_, err = db.CreateProfile(req.Username)
	if err != nil {
		http.Error(w, "handleCreateProfile(): FAILED", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "handleCreateProfile(): CREATED",
	})
}

func RegisterProfileRoutes(r *mux.Router) {
	profileRouter := r.PathPrefix("/api/profile").Subrouter()

	profileRouter.Use(AuthCheckerMiddleWare)
	profileRouter.HandleFunc("/me", handleGetProfile).Methods("GET")
	r.HandleFunc("/api/profile/new", handleCreateProfile).Methods("POST")
}
