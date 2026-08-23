package router

import (
	"golang-course-api/internal/controller"

	"github.com/gin-gonic/gin"
)

// New builds and returns the application router.
// router -> controller -> service -> repository -> models -> db
func New() *gin.Engine {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	v1 := r.Group("/api/v1")
	v1.GET("/ping", controller.NewPongController().Pong)
	v1.GET("/users", controller.NewUserController().GetUsers)

	return r
}
