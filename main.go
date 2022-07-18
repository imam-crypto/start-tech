package main

import (
	"fmt"
	"log"
	"pustaka-api/handler"
	"pustaka-api/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=root dbname=db_gorm port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected To db_gorm")

	// var users []user.User
	// length := len(users)
	// fmt.Println(length)

	// db.Find(&users)
	// length = len(users)
	// fmt.Println(length)

	// for _, user := range users {
	// 	fmt.Println(user.First_name)
	// 	fmt.Println(user.Email)
	// }

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	userByEmail, erors := userRepository.FindByEmail("response@gmail.com")
	if erors != nil {
		log.Fatal(erors.Error())

	}
	if userByEmail.ID == 0 {
		fmt.Println("user tidak di temuakan")
	} else {
		fmt.Println(userByEmail.First_name)
	}

	input := user.LoginInput{
		Email:    "response@gmail.com",
		Password: "response",
	}

	user, err := userService.Login(input)
	if err != nil {
		fmt.Println("terjadi kesalahan")
		fmt.Println(err, erors)
	}
	fmt.Print(user.Email)

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	router.Run()
}
