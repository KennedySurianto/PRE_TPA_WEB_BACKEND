package controller

import (
	"net/http"
	"pre_tpa_web/model/dto/request"
	"pre_tpa_web/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) CreateUser(context *gin.Context) {
	var userRequest request.UserRequest
	err := context.ShouldBindBodyWithJSON(&userRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	result, err := u.userService.CreateUser(userRequest.Name, userRequest.Email, userRequest.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"result": result})
}

func (u *UserController) GetAllUsers(c *gin.Context) {
	users, err := u.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (u *UserController) GetUser(c *gin.Context) {
	email := c.Param("email")
	user, err := u.userService.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (u *UserController) UpdateUser(c *gin.Context) {
	email := c.Param("email")
	name := c.PostForm("name")
	password := c.PostForm("password")

	err := u.userService.UpdateUser(email, name, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (u *UserController) DeleteUser(c *gin.Context) {
	email := c.Param("email")
	err := u.userService.DeleteUser(email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
