package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/malailiyati/beginnerBackend/internal/configs"
	"github.com/malailiyati/beginnerBackend/internal/routers"
)

func main() {
	// load .env
	_ = godotenv.Load()

	// init DB
	db, err := configs.InitDB()
	if err != nil {
		log.Println("Failed to connect DB:", err)
		return
	}
	defer db.Close()

	if err := configs.TestDB(db); err != nil {
		log.Println(" Ping DB failed:", err)
		return
	}
	log.Println("DB Connected")

	// init router
	r := routers.InitRouter(db)

	// run server di :8080
	r.Run(":8080")
}
