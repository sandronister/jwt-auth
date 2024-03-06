package userhandlers

import "github.com/sandronister/jwt-auth/internal/usecase/userusecase"

type RegisterInput struct {
	Username *string `json:"username" binding:"required"`
	Password *string `json:"password" binding:"required"`
}

type Handler struct {
	userUsecase *userusecase.User
}

func NewUserHandler(usecase *userusecase.User) *Handler {
	return &Handler{userUsecase: usecase}
}
