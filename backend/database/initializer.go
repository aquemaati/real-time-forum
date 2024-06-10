package database

import (
	"database/sql"
	"log"

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
	var err error

	Db, err = InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	//defer db.Close()

	if err := createTableUsers(Db); err != nil {
		log.Fatalf("Failed to create Users table: %v", err)
	}

	if err := createTableCategories(Db); err != nil {
		log.Fatalf("Failed to create Categories table : %v", err)
	}

	if err := createTablePosts(Db); err != nil {
		log.Fatalf("Failed to create Posts table: %v", err)
	}

	/*
	   	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
	   		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	   		defer cancel()

	   		if r.Method == http.MethodPost {
	   			nickname := r.FormValue("nickname")
	   			age := r.FormValue("age")
	   			gender := r.FormValue("gender")
	   			firstname := r.FormValue("firstname")
	   			lastname := r.FormValue("lastname")
	   			email := r.FormValue("email")
	   			password := r.FormValue("password")

	   			err := contexte.RegisterUser(ctx, db, nickname, age, gender, firstname, lastname, email, password)
	   			if err != nil {
	   				http.Error(w, err.Error(), http.StatusInternalServerError)
	   				return
	   			}

	   			fmt.Fprintf(w, "User %s successfully registered", nickname)
	   		} else {
	   			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	   		}
	   	})

	   	// Insert a test categories
	   , err = db.ExecContext(ctx, "INSERT OR IGNORE INTO Categories (name, description) VALUES (?, ?)",
	   		"This is a categorie", "This is a test description")
	   	if err != nil {
	   		log.Fatalf("Failed to insert test post: %v", err)
	   	}

	   	// Insert a test post
	   	_, err = db.ExecContext(ctx, `INSERT OR IGNORE INTO Posts (id, userId, title, description)
	   	VALUES (?, (SELECT id FROM Users WHERE nickname = ?), ?, ?)`, "1", "TEST", "This is a test post", "This is a test description")
	   	if err != nil {
	   		log.Fatalf("Failed to insert test post: %v", err)
	   	}*/
	log.Println("Database initialized, test user and test post inserted successfully")
}
