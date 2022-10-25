package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"start-tech/auth"
	"start-tech/campaign"
	"start-tech/handler"
	"start-tech/helper"
	"start-tech/payment"
	"start-tech/transaction"
	"start-tech/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	db_host := os.Getenv("DB_HOST")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	// db_tz := os.Getenv("DB_TZ")
	db_user := os.Getenv("DB_USER")
	port := os.Getenv("PORT")

	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		db_host, db_user, db_password, db_name, db_port,
	)

	// dsn := "host=localhost user=postgres password=root dbname=db_restbackend port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Connected To db_restbackend")

	paymentService := payment.NewServicePayment()

	authService := auth.NewService()
	campaignsRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	campaignService := campaign.NewService(campaignsRepository)
	transactionService := transaction.NewService(transactionRepository, campaignsRepository, paymentService)
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	er := db.AutoMigrate(user.User{}, campaign.CampaignImage{}, transaction.Transaction{})
	if er != nil {
		log.Fatal(er)
	}
	fmt.Println("Migrated")

	router := gin.Default()

	router.Static("/images", "user/images/")
	router.Use(cors.Default())
	// authService.GenerateToken(1001)

	api := router.Group("/api/v1")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvaibility)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/fetch", authMiddleware(authService, userService), userHandler.FetchUser)
	// route campaign
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns/create-campaign", authMiddleware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/update-campaign/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)

	// route for handler transaction

	api.GET("/campaigns/:id/transactions", transactionHandler.GetCampaignTransactions)
	api.GET("/campaigns/transactions-user", authMiddleware(authService, userService), transactionHandler.GeUserTransaction)
	api.POST("/transaction-create", authMiddleware(authService, userService), transactionHandler.Create)
	api.POST("/transactions/notification", transactionHandler.GetNotification)
	api.GET("/", campaignHandler.Tes)
	router.Run(port + ":")
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
