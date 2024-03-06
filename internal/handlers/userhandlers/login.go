package userhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Handler) Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := s.userUsecase.Login(*input.Username, *input.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
