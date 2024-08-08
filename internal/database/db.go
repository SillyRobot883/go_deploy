package database

import (
	"database/sql"
	"log"
	// the underscore is used to import the package without using it
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// InitDB initializes the SQLite database and creates the users table if it doesn't exist.
// this will be used in main.go
func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}

	// the above only creates a database, but we will also need to create the table
	// SQL statement to create the users table
	createTable := `CREATE TABLE IF NOT EXISTS users (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"username" TEXT NOT NULL UNIQUE,
		"email" TEXT NOT NULL UNIQUE,
		"password" TEXT NOT NULL,
		"created_at" DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	// Execute the SQL statement
	if _, err = DB.Exec(createTable); err != nil {
		log.Fatal(err)
	}
}
