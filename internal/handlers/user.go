package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sandronister/jwt-auth/internal/usecase"
)

type RegisterInput struct {
	Username *string `json:"username" binding:"required"`
	Password *string `json:"password" binding:"required"`
}

type UserHandler struct {
	userUsecase *usecase.User
}

func NewUserHandler(usecase *usecase.User) *UserHandler {
	return &UserHandler{userUsecase: usecase}
}

func (h *UserHandler) Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := h.userUsecase.CreateUser(*input.Username, *input.Password); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User created",
	})

}
