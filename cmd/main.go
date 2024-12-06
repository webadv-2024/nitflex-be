package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"nitflex/internal/handler"
	"nitflex/internal/repository"
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
	routes.Use(corsMiddleware())

	routes.POST("/register", h.Register)
	routes.POST("/login", h.Login)
	routes.POST("/login/google", h.GoogleLogin)

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
