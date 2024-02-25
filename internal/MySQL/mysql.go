
package mysql

import (
    "database/sql"
    "fmt"
    "log"
)

// Config holds MySQL configuration values.
type Config struct {
    Username string // MySQL username
    Password string // MySQL password
    Host     string // MySQL host
    Port     string // MySQL port
    Database string // MySQL database name
}

// Connect connects to a MySQL database using the provided configuration.
func Connect(config *Config) (*sql.DB, error) {
    // Compose the connection string using the configuration values
    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)

    // Open the database connection
    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        return nil, fmt.Errorf("error connecting to MySQL: %v", err)
    }

    // Check if the connection is established
    err = db.Ping()
    if err != nil {
        return nil, fmt.Errorf("error pinging MySQL: %v", err)
    }

    log.Println("Connected to MySQL")

    return db, nil
}

// InsertData inserts data into the MySQL database.
func InsertData(db *sql.DB, data string) error {
    // Example query to insert data into a table
    query := "INSERT INTO your_table_name (column_name) VALUES (?)"

    _, err := db.Exec(query, data)
    if err != nil {
        return fmt.Errorf("error inserting data into MySQL: %v", err)
    }

    log.Printf("Data inserted into MySQL: %s", data)

    return nil
}

// Close closes the database connection.
func Close(db *sql.DB) {
    err := db.Close()
    if err != nil {
        log.Printf("Error closing MySQL connection: %v\n", err)
    }
}
