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
	dsn := "host=localhost user=postgres password=root dbname=db_restbackend port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected To db_restbackend")

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
	userService.SaveAvatar(2, "/images/jpg")

	// er := db.AutoMigrate(user.User{})
	// if er != nil {
	// 	log.Fatal(er)
	// }
	// fmt.Println("Migrated")

	router := gin.Default()

	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvaibility)
	router.Run()
}
