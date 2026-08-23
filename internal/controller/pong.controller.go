package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(c *gin.Context) {
	// Return JSON response
	name := c.Param("name")
	name2 := c.Query("id")
	c.JSON(http.StatusFound, gin.H{
		"message": "pong " + name,
		"id":      name2,
	})
}
