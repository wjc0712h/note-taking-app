package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"note-taking-app/db"

	"github.com/gorilla/mux"
)

type User struct {
	Username string `json:"username"`
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Headers:", r.Header)
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	fmt.Println("Login attempt:", user.Username)
	_, err := db.GetProfile(user.Username)
	if err != nil {
		http.Error(w, "profile not found", http.StatusUnauthorized)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    user.Username,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteNoneMode,
	})
	fmt.Println("current user:", user.Username)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "handleLogin() successful")
}

func GetUsernameFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("username")
	if err != nil {
		fmt.Println("GetUsernameFromCookie() failed:", err)
		return "", fmt.Errorf("cookie not found")
	}
	return cookie.Value, nil
}

func AuthCheckerMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, err := GetUsernameFromCookie(r)

		if err != nil || username == "" {
			http.Error(w, "UNAUTHORIZED", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RegisterAuthRoutes(r *mux.Router) {
	authRouter := r.PathPrefix("/api/auth").Subrouter()
	authRouter.HandleFunc("/login", handleLogin).Methods("POST")
}
