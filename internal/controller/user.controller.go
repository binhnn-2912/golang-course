package controller

import (
	"golang-course-api/internal/service"
	"golang-course-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	name := c.DefaultQuery("name", "Hương")
	uid := c.Query("uid")

	response.SuccessResponse(c, 20001, gin.H{
		"message": uc.userService.GetUsersService(),
		"name":    name,
		"uid":     uid,
		"users":   []string{"Bình", "Trí", "Thường"},
	})
}
