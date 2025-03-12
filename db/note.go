package db

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

func CreateNote(content string, userName string) (Note, error) {
	newNote := Note{
		Id:        uuid.NewString(),
		Username:  userName,
		Content:   content,
		CreatedAt: time.Now().Format("2000-01-01 3:4:5 pm"),
	}

	query := `INSERT INTO notes (id, username, content, created_at) VALUES (?, ?, ?, ?)`
	_, err := DB.Exec(query, newNote.Id, newNote.Username, newNote.Content, newNote.CreatedAt)

	return newNote, err
}

func GetNote(id string) (Note, error) {
	query := `SELECT id, username, content, created_at FROM notes WHERE id=?`
	var note Note

	err := DB.QueryRow(query, id).Scan(&note.Id, &note.Username, &note.Content, &note.CreatedAt)
	return note, err
}

func GetNotesbyUsername(username string) ([]Note, error) {
	query := `SELECT id, username, content, created_at FROM notes WHERE username=?`

	rows, err := DB.Query(query, username)

	var notes []Note

	for rows.Next() {
		var note Note
		err = rows.Scan(&note.Id, &note.Username, &note.Content, &note.CreatedAt)

		notes = append(notes, note)
	}

	return notes, err
}
