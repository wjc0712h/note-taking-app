package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDBConn() {
	newDB := CreateDB()
	db, err := sql.Open("sqlite3", "./db/database.db")
	if err != nil {
		fmt.Println("connection failed:", err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("database.db is unreachable:", err)
		return
	}
	DB = db
	fmt.Println("connected to database.db successfully")

	if newDB {
		InitSchema()
	}
}

func CreateDB() bool {
	db := "./db/database.db"
	if _, err := os.Stat(db); os.IsNotExist(err) {
		fmt.Println("database.db not found, creating...")

		if err := os.MkdirAll("./db", os.ModePerm); err != nil {
			fmt.Println("failed to create directory:", err)
			return false
		}

		file, err := os.Create(db)
		if err != nil {
			fmt.Println("failed to create database.db:", err)
			return false
		}
		file.Close()
		fmt.Println("database.db created successfully")
		return true
	} else {
		fmt.Println("database.db already exists")
	}
	return false
}

func InitSchema() {
	schema := `
	CREATE TABLE IF NOT EXISTS profile (
		username TEXT NOT NULL PRIMARY KEY,
		createdAt TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS notes (
		id TEXT NOT NULL PRIMARY KEY,
		username TEXT NOT NULL,
		content TEXT NOT NULL,
		createdAt TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY(username) REFERENCES profile(username) ON DELETE CASCADE
	);
	`

	_, err := DB.Exec(schema)
	if err != nil {
		fmt.Println("failed to initialize database schema:", err)
	} else {
		fmt.Println("database.db schema initialized successfully")
	}
}
