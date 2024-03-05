package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/sandronister/jwt-auth/internal/di"
	"github.com/sandronister/jwt-auth/internal/infra/web"
)

func main() {
	server := web.NewWebServer(":8080")

	db, err := sql.Open("sqlite3", "./test.db")

	if err != nil {
		panic(db)
	}

	userHandler := di.NewUserHandler(db)
	server.AddRegisterHandler(&userHandler)

	server.Start()

}
