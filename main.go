package main

import (
	"fmt"
	"pre_tpa_web/controller"
	"pre_tpa_web/database"
	// "pre_tpa_web/model"
	"pre_tpa_web/repository"
	"pre_tpa_web/router"
	"pre_tpa_web/service"
)

func main() {
	fmt.Println("TPA WEB 25-1")

	studentName := "sigma"
	className := "TPA"
	className = "WEB" // overwrite class name

	fmt.Println(studentName)
	fmt.Println(className)

	classCodes := []string{"KN25-1", "OW25-1"}
	for _, code := range classCodes {
		fmt.Println(code)
	}

	userData := make(map[string]string)
	userData["username"] = "value1"
	userData["email"] = "value2"

	printMap(userData)

	delete(userData, "username") // remove the 'username' key

	printMap(userData)

	// connect db
	db := database.ConnectDatabase()

	// insert data to db
	userRepository := repository.NewUserRepository(db)
	// user := &model.User{
	// 	Name:     "John Doe",
	// 	Email:    "john.doe@gmail.com",
	// 	Password: "password123",
	// }
	// err := userRepository.CreateUser(user)
	// if err != nil {
	// 	fmt.Println("Error creating user:", err)
	// } else {
	// 	fmt.Println("User created successfully")
	// }

	// // create user using service
	userService := service.NewUserService(userRepository)
	// message, err := userService.CreateUser("Jane Doe", "jane.doe@gmail.com", "password456")
	// if err != nil {
	// 	fmt.Println("Error creating user:", err)
	// } else {
	// 	fmt.Println(message)
	// }

	userController := controller.NewUserController(userService)
	authService := service.NewAuthService(userRepository)
	authController := controller.NewAuthController(authService)

	r := &router.Router{
		UserController: userController,
		AuthController: authController,
	}

	router.NewRouter(r)
}

func printMap(data map[string]string) {
	fmt.Println("Printing map contents:")
	for key, val := range data {
		fmt.Println(key, val)
	}
}
