package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandronister/jwt-auth/config"
	"github.com/sandronister/jwt-auth/internal/di"
	"github.com/sandronister/jwt-auth/internal/infra/web"
)

func main() {
	err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	server := web.NewWebServer(fmt.Sprintf(":%s", os.Getenv("PORT")))

	db, err := sql.Open("sqlite3", "./test.db")

	if err != nil {
		panic(db)
	}

	userHandler := di.NewUserHandler(db)
	server.AddRegisterHandler(&userHandler)

	server.Start()

}
