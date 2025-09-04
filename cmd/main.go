package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/malailiyati/beginnerBackend/internal/configs"
	"github.com/malailiyati/beginnerBackend/internal/middlewares"
	"github.com/malailiyati/beginnerBackend/internal/routers"
)

func main() {
	// Load env (abaikan error kalau di prod sudah pakai env var)
	_ = godotenv.Load()

	// Init DB (pakai timeout biar fail-fast kalau salah kredensial)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := configs.InitPG(
		ctx,
		os.Getenv("DBUSER"),
		os.Getenv("DBPASS"),
		os.Getenv("DBHOST"),
		os.Getenv("DBPORT"),
		os.Getenv("DBNAME"),
	); err != nil {
		log.Fatalf("DB connect failed: %v", err)
	}
	defer configs.Pool.Close()
	log.Println("DB connected")

	// Init Gin + Middlewares (logger & CORS)
	r := gin.Default()
	r.Use(
		middlewares.RequestLogger(),
		middlewares.CORS(),
	)

	// Healthcheck simple
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Daftarkan semua route di satu tempat
	routers.Init(r)

	// 7) Run server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
