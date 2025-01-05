package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"nitflex/internal/handler"
	"nitflex/internal/middleware"
	"nitflex/internal/repository"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	mongoClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err = mongoClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	mongodb := mongoClient.Database("nitflex")
	h := handler.NewHandler(nil, mongodb)

	routes := gin.Default()

	// Add CORS middleware
	routes.Use(corsMiddleware())

	routes.POST("/register", h.Register)
	routes.POST("/login", h.Login)
	routes.POST("/login/google", h.GoogleLogin)

	routes.GET("/healthcheck", h.HealthCheck)
	routes.GET("/movies/trending", h.GetTrendingMovies)
	routes.GET("/movies", h.GetMovies)
	routes.GET("/movies/:id", h.GetMovieDetail)
	routes.GET("/casts", h.GetCasts)
	routes.GET("/casts/:id", h.GetCastInfo)
	routes.GET("/reviews/:id", h.GetReviews)
	routes.GET("/movies/popular", h.GetMoviePopular)
	routes.GET("/movies/upcoming", h.GetMovieUpcoming)

	// Add the middleware to protected routes
	routes.POST("/watchlist/", middleware.AuthMiddleware(), h.PostWatchlist)
	routes.GET("/watchlist/", middleware.AuthMiddleware(), h.GetWatchlist)
	routes.DELETE("/watchlist/", middleware.AuthMiddleware(), h.DeleteWatchlist)
	
	// Add favorite list routes
	routes.POST("/favorite/", middleware.AuthMiddleware(), h.PostFavoriteList)
	routes.GET("/favorite/", middleware.AuthMiddleware(), h.GetFavoriteList)
	routes.DELETE("/favorite/", middleware.AuthMiddleware(), h.DeleteFavoriteList)

	// Add rating routes
	routes.POST("/ratings/", middleware.AuthMiddleware(), h.PostRating)

	routes.Run(":3000")
}

func initDb() (*gorm.DB, error) {
	rawURL := os.Getenv("DATABASE_URL")
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		log.Fatalf("failed to parse database URL: %v", err)
		return nil, err
	}

	user := parsedURL.User.Username()
	password, _ := parsedURL.User.Password()
	host := parsedURL.Host
	path := strings.TrimPrefix(parsedURL.Path, "/")
	query := parsedURL.RawQuery

	// Ensure the query parameters are correctly formatted
	if query != "" {
		query = "?" + query
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s%s", user, password, host, path, query)

	// dsn := "root:123456@tcp(127.0.0.1:3343)/nitflex?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&repository.User{})
	return db, nil
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
