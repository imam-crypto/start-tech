package main

import (
	"fmt"
	"log"
	"net/http"
	"pustaka-api/auth"
	"pustaka-api/campaign"
	"pustaka-api/handler"
	"pustaka-api/helper"
	"pustaka-api/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	authService := auth.NewService()
	campaignsRepository := campaign.NewRepository(db)

	campaignService := campaign.NewService(campaignsRepository)

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	// er := db.AutoMigrate(campaign.Campaign{}, campaign.CampaignImage{})
	// if er != nil {
	// 	log.Fatal(er)
	// }
	// fmt.Println("Migrated")

	router := gin.Default()

	router.Static("/images", "/user/images")

	// authService.GenerateToken(1001)

	api := router.Group("/api/v1")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvaibility)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	// route campaign
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns/create-campaign", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	router.Run()
}
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("unauthorize", http.StatusUnauthorized, "failed", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("unauthorize", http.StatusUnauthorized, "failed", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("unauthorize", http.StatusUnauthorized, "failed", nil)

			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))
		user, err := userService.FindById(userID)
		if err != nil {
			response := helper.APIResponse("unauthorize", http.StatusUnauthorized, "failed", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("current_user", user)
	}

}
