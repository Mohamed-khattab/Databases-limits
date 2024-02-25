package databaseplayground

import (
	"databases-limit/configs"
	_ "databases-limit/internal/MySQL"
	"databases-limit/internal/PostgreSQL"
	"databases-limit/internal/SQLite"
	"fmt"
)

func main() {

	configs := configs.NewConfig()

	// Connect to PostgreSQL database

	db, err := postgres.Connect(configs)
	if err != nil {
		fmt.Println(err)
	}

	defer postgres.Close(db)

	// Insert a new record into the database in rotine

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			go postgres.Insert(db, "John Doe", 30)
		}
	}

}
