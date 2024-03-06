package di

import (
	"database/sql"

	"github.com/sandronister/jwt-auth/internal/handlers/userhandlers"
	"github.com/sandronister/jwt-auth/internal/infra/repositories"
	"github.com/sandronister/jwt-auth/internal/usecase/userusecase"
)

func NewUserHandler(conn *sql.DB) userhandlers.Handler {
	repository := repositories.NewUserRepository(conn)
	usecase := userusecase.NewUserUsecase(repository)
	return *userhandlers.NewUserHandler(usecase)
}
