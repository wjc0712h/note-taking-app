package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDBConn() {
	db, err := sql.Open("sqlite3", "./db/database.db")
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}

	// Check if the connection is actually working
	err = db.Ping()
	if err != nil {
		fmt.Println("Database is unreachable:", err)
		return
	}

	DB = db
	fmt.Println("Connected to SQLite successfully")
}
