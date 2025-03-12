package main

import (
	"github.com/gin-gonic/gin"
	"github.com/RedContritio/E-commerce-system/models"
	"github.com/RedContritio/E-commerce-system/services"
	"github.com/RedContritio/E-commerce-system/handlers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func initDB() *gorm.DB {
	dsn := "host=postgres user=postgres password=password dbname=ecommerce port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func main() {
	// 初始化数据库
	db := initDB()
	
	db.AutoMigrate(&models.User{})
	
	userService := services.UserService{DB: db}
	userHandler := handlers.UserHandler{UserService: &userService}
	
	// 初始化 Gin
	r := gin.Default()

	// 定义路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the E-commerce System!",
		})
	})

	r.POST("/register", userHandler.RegisterUser)

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}