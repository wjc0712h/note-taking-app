package db

import (
	"fmt"
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
		CreatedAt: time.Now().Format("2006/01/02, 3:04 pm"),
	}

	query := `INSERT INTO notes (id, username, content, createdAt) VALUES (?, ?, ?, ?)`
	_, err := DB.Exec(query, newNote.Id, newNote.Username, newNote.Content, newNote.CreatedAt)

	return newNote, err
}
func UpdateNote(id string, content string) (Note, error) {
	nowDateTimeStr := time.Now().Format("2006/01/02, 3:04 pm")
	note, _ := GetNote(id)

	if content != "" {
		note.Content = content
	}

	note.CreatedAt = nowDateTimeStr

	fmt.Println(note)
	query := `UPDATE notes SET content = ?, createdAt = ? WHERE id = ?`
	_, err := DB.Exec(query, note.Content, note.CreatedAt, note.Id)

	return note, err
}

func DeleteNote(id string) error {
	DB.Exec("PRAGMA foreign_keys = ON;")
	var count int
	err := DB.QueryRow("SELECT COUNT(*) FROM notes WHERE id = ?", id).Scan(&count)
	if err != nil {
		fmt.Println("error checking note existence:")
		return fmt.Errorf("error checking note existence: %v", err)
	}

	if count == 0 {
		fmt.Printf("note with id %s does not exist", id)

		return fmt.Errorf("note with id %s does not exist", id)
	}

	// Now delete the note
	query := `DELETE FROM notes WHERE id = ?`
	result, err := DB.Exec(query, id)
	if err != nil {
		fmt.Println("error executing delete query:")
		return fmt.Errorf("error executing delete query: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("error getting affected rows")
		return fmt.Errorf("error getting affected rows: %v", err)
	}

	if rowsAffected == 0 {
		fmt.Println("note was not deleted")
		return fmt.Errorf("note was not deleted")
	}
	fmt.Println("deleted")
	return nil
}

func GetNote(id string) (Note, error) {
	query := `SELECT id, username, content, createdAt FROM notes WHERE id=?`
	var note Note

	err := DB.QueryRow(query, id).Scan(&note.Id, &note.Username, &note.Content, &note.CreatedAt)
	return note, err
}

func GetNotesbyUsername(username string) ([]Note, error) {
	query := `SELECT id, username, content, createdAt FROM notes WHERE username=?`

	rows, err := DB.Query(query, username)

	var notes []Note

	for rows.Next() {
		var note Note
		err = rows.Scan(&note.Id, &note.Username, &note.Content, &note.CreatedAt)

		notes = append(notes, note)
	}

	return notes, err
}
