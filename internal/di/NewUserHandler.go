package di

import (
	"database/sql"

	"github.com/sandronister/jwt-auth/internal/handlers"
	"github.com/sandronister/jwt-auth/internal/infra/repositories"
	"github.com/sandronister/jwt-auth/internal/usecase"
)

func NewUserHandler(conn *sql.DB) handlers.UserHandler {
	repository := repositories.NewUserRepository(conn)
	usecase := usecase.NewUserUsecase(repository)
	return *handlers.NewUserHandler(usecase)
}
