package userhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Current(c *gin.Context) {

	id, _ := c.Get("user_id")

	user, err := h.userUsecase.GetUserById(id.(int))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
