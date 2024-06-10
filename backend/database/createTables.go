package database

import (
	"database/sql"
	"fmt"
)

// createTableUsers creates the Users table
func createTableUsers(db *sql.DB) error {
	table := `
    CREATE TABLE IF NOT EXISTS Users (
        id TEXT NOT NULL,
        nickname TEXT NOT NULL UNIQUE,
        age INTEGER,
        gender TEXT,
        firstname TEXT,
        lastname TEXT,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL,
        created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		image TEXT NOT NULL DEFAULT 'backend/database/img/img.png'
    );
    `
	_, err := db.Exec(table)
	return err
}

// createTableCategories creates Categories table
func createTableCategories(db *sql.DB) error {
	table := `
	CREATE TABLE IF NOT EXISTS Categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT 
	);
	`
	_, err := db.Exec(table)
	return err
}

// createTablePosts creates the Posts table
func createTablePosts(db *sql.DB) error {
	table := `
	CREATE TABLE IF NOT EXISTS Posts (
		id TEXT NOT NULL,
		userId TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		FOREIGN KEY(userId) REFERENCES Users(id)
	);`
	fmt.Println("HERE")
	_, err := db.Exec(table)
	return err
}
