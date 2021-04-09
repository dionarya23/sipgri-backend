package main

import (
	"log"
	"os"

	"github.com/dionarya23/sipgri-backend/auth"
	"github.com/dionarya23/sipgri-backend/guru"
	"github.com/dionarya23/sipgri-backend/handlers"
	"github.com/dionarya23/sipgri-backend/mata_pelajaran"
	"github.com/dionarya23/sipgri-backend/middleware"
	"github.com/dionarya23/sipgri-backend/peserta_didik"
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
	pesertaDidikRepository := peserta_didik.NewRepository(db)

	guruService := guru.NewService(guruRepository)
	authService := auth.NewService()
	mataPelajaranService := mata_pelajaran.NewService(mataPelajaranRepository)
	pesertaDidikService := peserta_didik.NewService(pesertaDidikRepository)

	guruHandler := handlers.NewGuruHandler(guruService, authService)
	mataPelajaranHandler := handlers.NewMataPelajaranHandler(mataPelajaranService)
	pesertaDidikHandler := handlers.NewPesertaDidikHandler(pesertaDidikService)

	router := gin.Default()

	apiHandler := router.Group("/api")

	apiHandler.POST("/guru/register", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), guruHandler.RegisterGuru)
	apiHandler.POST("/guru/login", guruHandler.Login)
	apiHandler.GET("/guru/", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), guruHandler.GetAllGuru)
	apiHandler.GET("/guru/:nip", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru"}), guruHandler.GetOneGuru)
	apiHandler.PUT("/guru/:nip", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru"}), guruHandler.UpdateGuru)
	apiHandler.DELETE("/guru/:nip", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), guruHandler.DeleteGuru)
	apiHandler.POST("/guru/check-nip", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru"}), guruHandler.IsNipExist)
	apiHandler.POST("/guru/check-email", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru"}), guruHandler.IsEmailExist)

	apiHandler.POST("/mata-pelajaran/", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mataPelajaranHandler.CreateNewMataPelajaran)
	apiHandler.GET("/mata-pelajaran/", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mataPelajaranHandler.GetAll)
	apiHandler.GET("/mata-pelajaran/:id_mata_pelajaran", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mataPelajaranHandler.GetOne)
	apiHandler.PUT("/mata-pelajaran/:id_mata_pelajaran", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mataPelajaranHandler.UpdatedMataPelajaran)
	apiHandler.DELETE("/mata-pelajaran/:id_mata_pelajaran", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mataPelajaranHandler.DeleteByIDMataPelajaran)

	apiHandler.POST("/peserta-didik/", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), pesertaDidikHandler.CreatePesertaDidik)
	apiHandler.GET("/peserta-didik/", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), pesertaDidikHandler.GetAllPesertaDidik)
	apiHandler.GET("/peserta-didik/one/", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru"}), pesertaDidikHandler.GetOnePesertaDidik)
	apiHandler.PUT("/peserta-didik/:nisn/", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), pesertaDidikHandler.UpdatePesertaDidik)
	apiHandler.DELETE("/peserta-didik/:nisn/", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), pesertaDidikHandler.DeleteByNisn)

	router.Run()
}
