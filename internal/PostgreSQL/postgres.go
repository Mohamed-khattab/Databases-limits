package postgres

import (
	"database/sql"
	"databases-limit/configs"
	"fmt"
	"log"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// Connect to PostgreSQL database
func Connect(config *configs.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost, config.DbPort, config.DbUser, config.DbPass, config.DbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("Error connecting to database: %v", err)
	}
	// Check if the database is alive
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf("Error pinging database: %v", err)
	}

	log.Println("Connected to database")
	return db, nil
}

// Insert a new record into the database
func Insert(db *sql.DB, name string, age int) error {
	_, err := db.Exec("INSERT INTO users(name, age) VALUES($1, $2)", name, age)
	if err != nil {
		return fmt.Errorf("error inserting record: %v", err)
	}
	log.Println("Inserted a new record")
	return nil
}

// Close the database connection
func Close(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
