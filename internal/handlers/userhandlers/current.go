package userhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandronister/jwt-auth/internal/tools/tokenservice"
)

func (h *Handler) Current(c *gin.Context) {
	user_id, err := tokenservice.ExtractTokenMetadata(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userUsecase.GetUserById(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
