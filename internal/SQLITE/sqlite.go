// internal/db/sqlite/sqlite.go

package sqlite

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3" // SQLite driver
)

// Connect connects to an SQLite database using the provided connection string.
func Connect(connectionString string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", connectionString)
    if err != nil {
        return nil, fmt.Errorf("error connecting to SQLite: %v", err)
    }

    // Check if the connection is established
    err = db.Ping()
    if err != nil {
        return nil, fmt.Errorf("error pinging SQLite: %v", err)
    }

    log.Println("Connected to SQLite")

    return db, nil
}

// InsertData inserts data into the SQLite database.
func InsertData(db *sql.DB, data string) error {
    // Example query to insert data into a table
    query := "INSERT INTO your_table_name (column_name) VALUES (?)"

    _, err := db.Exec(query, data)
    if err != nil {
        return fmt.Errorf("error inserting data into SQLite: %v", err)
    }

    log.Printf("Data inserted into SQLite: %s", data)

    return nil
}

// Close closes the database connection.
func Close(db *sql.DB) {
    err := db.Close()
    if err != nil {
        log.Printf("Error closing SQLite connection: %v\n", err)
    }
}
