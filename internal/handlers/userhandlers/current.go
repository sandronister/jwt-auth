package userhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) Current(c *gin.Context) {

	id, _ := c.Get("user_id")

	user, err := h.userUsecase.GetUserById(id.(int))

	userDTO := UserDTO{
		ID:       user.ID,
		Username: user.Username,
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": userDTO})
}
