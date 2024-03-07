package userhandlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Current(c *gin.Context) {

	user_id, exists := c.Get("user_id")
	fmt.Println(user_id)

	if exists == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}

	user, err := h.userUsecase.GetUserById(user_id.(int))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
