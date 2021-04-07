package main

import (
	"log"
	"os"

	"github.com/dionarya23/sipgri-backend/guru"
	"github.com/dionarya23/sipgri-backend/handlers"
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

	guruHandler := handlers.NewGuruHandler(guruService)

	router := gin.Default()

	api := router.Group("/api")

	api.POST("/guru/register", guruHandler.RegisterGuru)

	router.Run()
}
