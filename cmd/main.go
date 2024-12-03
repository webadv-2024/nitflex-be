package main

import (
	"fmt"
	"log"
	"nitflex/internal/handler"
	"nitflex/internal/repository"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db, err := initDb()
	if err != nil {
		panic("failed to connect to database")
	}

	h := handler.NewHandler(db)

	routes := gin.Default()

	// Add CORS middleware
	routes.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("WEB_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.POST("/register", h.Register)
	routes.POST("/login", h.Login)
	routes.GET("/auth/google", h.GoogleLogin)
	routes.GET("/auth/google/callback", h.GoogleCallback)

	routes.Run(":3000")
}

func initDb() (*gorm.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3343)/nitflex?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&repository.User{})
	return db, nil
}
