package main

import (
	"fmt"
	"log"
	"time"

	"bone-backend/db"
	"bone-backend/handlers"
	"bone-backend/middleware"
	"bone-backend/models"

	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Database configuration
	dbConfig := db.Config{
		Host:     os.Getenv("POSTGRESQLHOST"),
		Port:     os.Getenv("POSTGRESQLDATABASEPORT"),
		User:     os.Getenv("POSTGRESQLUSER"),
		Password: os.Getenv("POSTGRESQLPASSWORD"),
		DBName:   os.Getenv("POSTGRESQLDATABASENAME"),
		SSLMode:  "disable", // Use "disable" for development and "require" for production
	}

	// Establish database connection
	database, err := db.NewConnection(dbConfig)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	database.AutoMigrate(&models.User{}, &models.Article{}, &models.ThreadTag{}, &models.Thread{}, &models.ThreadThreadTag{}, &models.ThreadLike{})

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // Allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // Allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
		AllowCredentials: true,                                                // Allow cookies to be sent with requests
		ExposeHeaders:    []string{"X-Custom-Header"},                         // Expose additional headers
		MaxAge:           12 * time.Hour,                                      // Cache preflight response for 12 hours
	}))

	// ROUTES
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/users/signup", func(c *gin.Context) {
		handlers.Signup(database, c.Writer, c.Request)
	})

	r.POST("/api/users/signin", func(c *gin.Context) {
		handlers.Signin(database, c.Writer, c.Request)
	})

	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/user-data", func(c *gin.Context) {
		userID, err := uuid.Parse(c.MustGet("user_id").(string))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		handlers.GetUserByID(database, c.Writer, c.Request, userID)
	})

	protected.GET("/users", func(c *gin.Context) {
		handlers.GetUsers(database, c.Writer, c.Request)
	})

	protected.POST("/users/update-last-questionnaire-date", func(c *gin.Context) {
		userID, err := uuid.Parse(c.MustGet("user_id").(string))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		handlers.UpdateUserLastQuestionnaireDate(database, c.Writer, c.Request, userID)
	})

	protected.GET("/articles", func(c *gin.Context) {
		handlers.GetAllArticle(database, c.Writer, c.Request)
	})

	protected.GET("/articles/latest", func(c *gin.Context) {
		handlers.GetLatestArticle(database, c.Writer, c.Request)
	})

	protected.GET("/threads", func(c *gin.Context) {
		userID, err := uuid.Parse(c.MustGet("user_id").(string))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		handlers.GetAllThread(database, c.Writer, c.Request, userID)
	})

	protected.POST("/threads/like", func(c *gin.Context) {
		userID, err := uuid.Parse(c.MustGet("user_id").(string))
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid user ID"})
			return
		}
		handlers.ThreadLike(database, c.Writer, c.Request, userID)
	})

	// Run the server

	port := fmt.Sprintf(":%s", os.Getenv("BACKENDPORT"))
	r.Run(port)
}
