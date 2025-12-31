package main

import (
	"context"
	"flag"
	"log"
	"mampu-wallet/database"
	"mampu-wallet/docs"
	"mampu-wallet/internal/router"
	"mampu-wallet/internal/tools"
	"os"

	"github.com/gin-gonic/gin"
)

// @title           API Wallet MAMPU.IO
// @version         1.0
// @description     Ini adalah dokumentasi API untuk sistem manajemen wallet.
// @termsOfService  http://swagger.io/terms/
// @BasePath /api/v1
func main() {
	tools.LoadEnv()

	migrate := flag.Bool("m", false, "Jalankan migrasi tabel")
	seed := flag.Bool("s", false, "Jalankan pengisian data awal (seeding)")
	flag.Parse()

	db := database.InitDB()
	if db == nil {
		log.Fatal("database not connected.")
		os.Exit(1)
	}

	if *migrate {
		database.RunMigration(db)
	}

	if *seed {
		database.SeedData(db)
	}

	docs.SwaggerInfo.BasePath = "/" + os.Getenv("PREFIX_API")

	r := gin.Default()
	ctx := context.Background()

	rh := &router.Handler{
		DB:  db,
		R:   r,
		CTX: ctx,
	}

	rh.Routes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}

	r.Run(":" + port)
}
