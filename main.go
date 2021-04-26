package main

import (
	"log"
	"os"

	"github.com/dionarya23/sipgri-backend/auth"
	"github.com/dionarya23/sipgri-backend/estrakulikuler"
	"github.com/dionarya23/sipgri-backend/guru"
	"github.com/dionarya23/sipgri-backend/handlers"
	"github.com/dionarya23/sipgri-backend/jadwal"
	"github.com/dionarya23/sipgri-backend/kelas"
	"github.com/dionarya23/sipgri-backend/mata_pelajaran"
	"github.com/dionarya23/sipgri-backend/mengajar"
	"github.com/dionarya23/sipgri-backend/middleware"
	"github.com/dionarya23/sipgri-backend/peserta_didik"
	"github.com/gin-contrib/cors"
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
	kelasRepository := kelas.NewRepository(db)
	mengajarRepository := mengajar.NewRepository(db)
	eskulRepository := estrakulikuler.NewRepository(db)
	jadwalRepository := jadwal.NewRepository(db)

	guruService := guru.NewService(guruRepository)
	authService := auth.NewService()
	mataPelajaranService := mata_pelajaran.NewService(mataPelajaranRepository)
	pesertaDidikService := peserta_didik.NewService(pesertaDidikRepository)
	kelasService := kelas.NewService(kelasRepository)
	mengajarService := mengajar.NewService(mengajarRepository)
	eskulService := estrakulikuler.NewService(eskulRepository)
	jadwalService := jadwal.NewService(jadwalRepository)

	guruHandler := handlers.NewGuruHandler(guruService, authService)
	mataPelajaranHandler := handlers.NewMataPelajaranHandler(mataPelajaranService)
	pesertaDidikHandler := handlers.NewPesertaDidikHandler(pesertaDidikService)
	kelasHandler := handlers.NewKelasHandler(kelasService)
	mengajarHandler := handlers.NewMengajarHandler(mengajarService)
	eskulHandler := handlers.NewEskulHandler(eskulService)
	jadwalHandler := handlers.NewJadwalHandler(jadwalService)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  false,
		AllowOriginFunc:  func(origin string) bool { return true },
		MaxAge:           86400,
	}))

	apiHandler := router.Group("/api")

	apiHandler.POST("/guru/register", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), guruHandler.RegisterGuru)
	apiHandler.POST("/guru/login", guruHandler.Login)
	apiHandler.GET("/guru/", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), guruHandler.GetAllGuru)
	apiHandler.GET("/guru/:nip", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru"}), guruHandler.GetOneGuru)
	apiHandler.PUT("/guru/:nip", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru"}), guruHandler.UpdateGuru)
	apiHandler.DELETE("/guru/:nip", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), guruHandler.DeleteGuru)
	apiHandler.POST("/guru/check-nip", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru"}), guruHandler.IsNipExist)
	apiHandler.POST("/guru/check-email", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru"}), guruHandler.IsEmailExist)

	apiHandler.POST("/mata-pelajaran", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mataPelajaranHandler.CreateNewMataPelajaran)
	apiHandler.GET("/mata-pelajaran", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mataPelajaranHandler.GetAll)
	apiHandler.GET("/mata-pelajaran/:id_mata_pelajaran", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mataPelajaranHandler.GetOne)
	apiHandler.PUT("/mata-pelajaran/:id_mata_pelajaran", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mataPelajaranHandler.UpdatedMataPelajaran)
	apiHandler.DELETE("/mata-pelajaran/:id_mata_pelajaran", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mataPelajaranHandler.DeleteByIDMataPelajaran)

	apiHandler.POST("/peserta-didik", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), pesertaDidikHandler.CreatePesertaDidik)
	apiHandler.GET("/peserta-didik", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), pesertaDidikHandler.GetAllPesertaDidik)
	apiHandler.GET("/peserta-didik/one", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru"}), pesertaDidikHandler.GetOnePesertaDidik)
	apiHandler.PUT("/peserta-didik/:nisn", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), pesertaDidikHandler.UpdatePesertaDidik)
	apiHandler.DELETE("/peserta-didik/:nisn", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), pesertaDidikHandler.DeleteByNisn)
	apiHandler.GET("/peserta-didik/wali-kelas/:nip_wali", middleware.AuthMiddleware(authService, guruService, []string{"admin", "wali_kelas"}), kelasHandler.GetByNipWali)

	apiHandler.POST("/kelas", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), kelasHandler.CreateKelas)
	apiHandler.GET("/kelas", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), kelasHandler.GetAll)
	apiHandler.GET("/kelas/:id_kelas", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), kelasHandler.GetById)
	apiHandler.PUT("/kelas/:id_kelas", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), kelasHandler.UpdateById)
	apiHandler.DELETE("/kelas/:id_kelas", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), kelasHandler.DeleteById)

	apiHandler.GET("/mengajar/guru/:nip_guru", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru", "wali_kelas"}), mengajarHandler.GetByNipGuru)
	apiHandler.POST("/mengajar", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mengajarHandler.Create)
	apiHandler.GET("/mengajar", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mengajarHandler.GetAll)
	apiHandler.GET("/mengajar/one/:kode_mengajar", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mengajarHandler.GetByKodeMengajar)
	apiHandler.PUT("/mengajar/:kode_mengajar", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mengajarHandler.UpdateMengajar)
	apiHandler.DELETE("/mengajar/:kode_mengajar", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), mengajarHandler.DeleteByKodeMengajar)

	apiHandler.POST("/eskul", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), eskulHandler.Create)
	apiHandler.GET("/eskul", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), eskulHandler.GetAll)
	apiHandler.GET("/eskul/one/:id_estrakulikuler", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), eskulHandler.GetById)
	apiHandler.GET("/eskul/pembimbing/:nip_pembimbing", middleware.AuthMiddleware(authService, guruService, []string{"admin", "guru", "wali_kelas"}), eskulHandler.GetByNipGuru)
	apiHandler.PUT("/eskul/:id_estrakulikuler", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), eskulHandler.UpdateById)
	apiHandler.DELETE("/eskul/:id_estrakulikuler", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), eskulHandler.DeleteById)

	apiHandler.POST("/jadwal", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), jadwalHandler.CreateNewData)
	apiHandler.GET("/jadwal", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), jadwalHandler.FindAllJadwal)
	apiHandler.GET("/jadwal/one/:id_jadwal", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), jadwalHandler.FindOneByIdJadwal)
	apiHandler.PUT("/jadwal/:id_jadwal", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), jadwalHandler.UpdateById)
	apiHandler.DELETE("/jadwal/:id_jadwal", middleware.AuthMiddleware(authService, guruService, []string{"admin"}), jadwalHandler.DeleteById)

	router.Run()
}
