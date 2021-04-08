package main

import (
	"log"
	"os"

	"github.com/dionarya23/sipgri-backend/auth"
	"github.com/dionarya23/sipgri-backend/guru"
	"github.com/dionarya23/sipgri-backend/handlers"
	"github.com/dionarya23/sipgri-backend/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	guruRepository := guru.NewRepository(db)

	guruService := guru.NewService(guruRepository)
	authService := auth.NewService()

	guruHandler := handlers.NewGuruHandler(guruService, authService)

	router := gin.Default()

	apiGuru := router.Group("/api/guru")
	apiGuru.POST("/register", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), guruHandler.RegisterGuru)
	apiGuru.POST("/login", guruHandler.Login)
	apiGuru.GET("/", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), guruHandler.GetAllGuru)
	apiGuru.GET("/:nip", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru"}), guruHandler.GetOneGuru)
	apiGuru.PUT("/:nip", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru"}), guruHandler.UpdateGuru)
	apiGuru.DELETE("/:nip", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), guruHandler.DeleteGuru)

	router.Run()
}
