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
