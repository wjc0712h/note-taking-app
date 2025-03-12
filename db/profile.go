package db

import (
	"time"
)

type Profile struct {
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
}

func CreateProfile(username string) (Profile, error) {
	newProfile := Profile{
		Username:  username,
		CreatedAt: time.Now().Format("2000-01-01 3:4:5 pm"),
	}
	query := `INSERT INTO profile (username, created_at) VALUES (?, ?)`
	_, err := DB.Exec(query, newProfile.Username, newProfile.CreatedAt)

	return newProfile, err
}

func GetProfile(username string) (Profile, error) {
	query := `SELECT username, created_at FROM profile WHERE username=?`
	var user Profile
	err := DB.QueryRow(query, username).Scan(&user.Username, &user.CreatedAt)
	return user, err
}
