package store

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jl-calda/geth-starter/config"
)

func InitDB() {
	// Initialize

	dbConfig := config.LoadDBConfig()
	url, err := dbConfig.GetURL()

	if err != nil {
		fmt.Fprintf(os.Stderr, "No DB url given. Set ENV var or .env file %d\n", err)
		os.Exit(1)
	}

	conn, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s\n", url, err)
		os.Exit(1)
	}
	defer conn.Close()

}

// func createTables() error {
// 	createUserTable := `
// 		CREATE TABLE IF NOT EXISTS users (
// 			pk INTEGER PRIMARY KEY AUTOINCREMENT,
// 			id VARCHAR(36) DEFAULT(UUID()),
// 			email STRING

// 		)`

// }
