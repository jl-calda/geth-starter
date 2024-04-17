package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DB           *sql.DB
	InProduction bool
}

func (a *AppConfig) SetDb(db *sql.DB) {
	a.DB = db
}

func (a *AppConfig) GetEnvironment() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Fprintf(os.Stdout, ".env not found: %s\n", err)
	}

	env := os.Getenv("ENVIRONMENT")

	var inProduction bool

	if env == "production" {
		inProduction = true
	} else {
		inProduction = false
	}

	a.InProduction = inProduction
}
