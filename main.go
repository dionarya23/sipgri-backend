package main

import (
	"log"
	"os"

	"github.com/dionarya23/sipgri-backend/auth"
	"github.com/dionarya23/sipgri-backend/guru"
	"github.com/dionarya23/sipgri-backend/handlers"
	"github.com/dionarya23/sipgri-backend/mata_pelajaran"
	"github.com/dionarya23/sipgri-backend/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	guruRepository := guru.NewRepository(db)
	mataPelajaranRepository := mata_pelajaran.NewRepository(db)

	guruService := guru.NewService(guruRepository)
	authService := auth.NewService()
	mataPelajaranService := mata_pelajaran.NewService(mataPelajaranRepository)

	guruHandler := handlers.NewGuruHandler(guruService, authService)
	mataPelajaranHandler := handlers.NewMataPelajaranHandler(mataPelajaranService)

	router := gin.Default()

	apiGuru := router.Group("/api/guru")
	apiGuru.POST("/register", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), guruHandler.RegisterGuru)
	apiGuru.POST("/login", guruHandler.Login)
	apiGuru.GET("/", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), guruHandler.GetAllGuru)
	apiGuru.GET("/:nip", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru"}), guruHandler.GetOneGuru)
	apiGuru.PUT("/:nip", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru"}), guruHandler.UpdateGuru)
	apiGuru.DELETE("/:nip", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), guruHandler.DeleteGuru)
	apiGuru.POST("/check-nip", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru"}), guruHandler.IsNipExist)
	apiGuru.POST("/check-email", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru"}), guruHandler.IsEmailExist)

	apiMataPelajaran := router.Group("api/mata-pelajaran")
	apiMataPelajaran.POST("/", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mataPelajaranHandler.CreateNewMataPelajaran)
	apiMataPelajaran.GET("/", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mataPelajaranHandler.GetAll)
	apiMataPelajaran.GET("/:id_mata_pelajaran", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mataPelajaranHandler.GetOne)
	apiMataPelajaran.PUT("/:id_mata_pelajaran", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mataPelajaranHandler.UpdatedMataPelajaran)
	apiMataPelajaran.DELETE("/:id_mata_pelajaran", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mataPelajaranHandler.DeleteByIDMataPelajaran)

	router.Run()
}
