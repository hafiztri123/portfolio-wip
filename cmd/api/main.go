package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hafiztri123/internal/database"
	"github.com/hafiztri123/internal/handler"
	"github.com/hafiztri123/internal/models"
	"github.com/hafiztri123/internal/repository"
	"github.com/hafiztri123/internal/utils"
	"github.com/joho/godotenv"
)

const (
    BASE_API_URL = "/api/v1"
    AUTH_URL = BASE_API_URL + "/auth"
    PORTFOLIO_URL = BASE_API_URL + "/portfolio"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file", err)
		return 
	}

    router := gin.Default()
    
    // CORS middleware
    config := cors.DefaultConfig()
    config.AllowOrigins = []string{"http://localhost:5173"}
    config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
    config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
    config.AllowCredentials = true
    config.MaxAge = 12 * time.Hour
    router.Use(cors.New(config))

	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}

    //Auth
	authRepositry := repository.NewAuthRepository(db)
    authHandler := handler.NewAuthHandler(authRepositry)
    //Portfolio
    portfolioRepository := repository.NewPortfolioRepository(db)
    portfolioHandler := handler.NewPortfolioHandler(portfolioRepository)
    


	auth := router.Group(BASE_API_URL)
	{
		auth.POST("/register", authHandler.HandleRegister)
		auth.POST("/login", authHandler.HandleLogin)
	}

    portfolio := router.Group(PORTFOLIO_URL)
    portfolio.Use(utils.AuthMiddleware())
    {
        portfolio.GET("/", portfolioHandler.GetAllPortfolios)
        portfolio.POST("/", portfolioHandler.CreatePortfolio)
    }



	router.Run(":8080")

}


