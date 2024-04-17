package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/jl-calda/geth-starter/config"
	"github.com/jl-calda/geth-starter/handlers"

	"github.com/labstack/echo/v4"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

var appConfig config.AppConfig

func main() {
	app := echo.New()
	app.Static("/static", "static")

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
	appConfig.SetDb(conn)
	defer conn.Close()

	app.GET("/", handlers.HomePageHandler)
	app.GET("/login", handlers.LoginPageHandler)

	app.Start("127.0.0.1:8080")
}
