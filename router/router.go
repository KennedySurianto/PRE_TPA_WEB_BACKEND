package router

import (
	"pre_tpa_web/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Router struct {
	UserController *controller.UserController
	AuthController *controller.AuthController
}

func NewRouter(r *Router) {
	// Initialize the Gin router
	router := gin.Default()
	router.Use(cors.Default());

	// User routes
	router.POST("/", r.UserController.CreateUser)
	router.GET("/users", r.UserController.GetAllUsers)
	router.GET("/users/:email", r.UserController.GetUser)
	router.PUT("/users/:email", r.UserController.UpdateUser)
	router.DELETE("/users/:email", r.UserController.DeleteUser)

	// Auth Routes
	router.POST("/auth/register", r.AuthController.Register)
	router.POST("/auth/login", r.AuthController.Login)
	router.POST("/auth/logout", r.AuthController.Logout)

	router.Run()
}
