package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	URL   string
	Token string
}

func LoadDBConfig() DBConfig {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Fprintf(os.Stdout, ".env not found: %s\n", err)
	}

	url := os.Getenv("DB_URL")
	token := os.Getenv("DB_TOKEN")

	return DBConfig{URL: url, Token: token}
}

func (c *DBConfig) GetURL() (string, error) {
	url := c.URL

	if len(c.Token) > 0 {
		url = url + "?authToken=" + c.Token
	}

	if len(url) == 0 {
		return "", errors.New("no database url given")
	}
	return url, nil

}
