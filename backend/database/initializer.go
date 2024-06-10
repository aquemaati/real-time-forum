package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// InitDB initializes the connection to the database
func InitDB() (*sql.DB, error) {
	dbFilePath := "./backend/database/database.db"

	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return nil, err
	}

	// Configure database connection settings
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(0)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// InitMainDB initializes the main database
func InitMainDB() {
	db, err := InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer db.Close()

	if err := createTableUsers(db); err != nil {
		log.Fatalf("Failed to create Users table: %v", err)
	}

	if err := createTableCategories(db); err != nil {
		log.Fatalf("Failed to create Categories table : %v", err)
	}

	if err := createTablePosts(db); err != nil {
		log.Fatalf("Failed to create Posts table: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Insert a test user
	_, err = db.ExecContext(ctx, "INSERT OR IGNORE INTO Users (id, nickname, age, gender, firstname, lastname, email, password) VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		"NOT NULL", "TEST", 18, "male", "firstname", "lastname", "test@test.test", "qwerty")
	if err != nil {
		log.Fatalf("Failed to insert test user: %v", err)
	}

	// Insert a test categories
	_, err = db.ExecContext(ctx, "INSERT OR IGNORE INTO Categories (name, description) VALUES (?, ?)",
		"This is a categorie", "This is a test description")
	if err != nil {
		log.Fatalf("Failed to insert test post: %v", err)
	}

	// Insert a test post
	_, err = db.ExecContext(ctx, `INSERT OR IGNORE INTO Posts (id, userId, title, description) 
	VALUES (?, (SELECT id FROM Users WHERE nickname = ?), ?, ?)`, "1", "TEST", "This is a test post", "This is a test description")
	if err != nil {
		log.Fatalf("Failed to insert test post: %v", err)
	}
	log.Println("Database initialized, test user and test post inserted successfully")
}
