package userhandlers

import "github.com/gin-gonic/gin"

func (h *Handler) Register(c *gin.Context) {
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
