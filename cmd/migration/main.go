package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "modernc.org/sqlite"
)

// applyMigration reads and executes the SQL migration file
func applyMigration(db *sql.DB, filename string) error {
	// Read SQL file
	query, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read migration file: %w", err)
	}

	// Execute migration SQL
	_, err = db.Exec(string(query))
	if err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	log.Println("Migration applied successfully!")
	return nil
}

func main() {
	// Define database file name
	dbPath := "./store/database.db"
	migrationFile := "./cmd/migration/migration.sql" // Ensure this file contains the SQL script

	// Open SQLite database
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	// Apply migration
	if err := applyMigration(db, migrationFile); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	InsertUser(db)
}

func InsertUser(db *sql.DB) error {
	insertQuery := `INSERT INTO Users (userId, tenantId, email) VALUES (?, ?, ?);`
	_, err := db.Exec(insertQuery, "asfar.sharief", "company1", "asfarsharief015@gmail.com")
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}
