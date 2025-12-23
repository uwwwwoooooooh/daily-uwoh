package main

import (
	"log"

	"github.com/uwwwwoooooooh/daily-uwoh/internal/api/handler"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/api/router"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/config"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/db"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/repository"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/service"
)

func main() {
	log.Println("Initializing DailyUwoh System...")

	cfg := config.LoadConfig()
	log.Printf("Loaded config. Port: %s", cfg.ServerPort)

	if cfg.JWTSecret == "secret" {
		log.Println("WARNING: Using default JWT secret 'secret'. This is insecure for production!")
	}

	db, err := db.ConnectDB(cfg.DBUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	store := repository.NewStore(db)
	authService := service.NewAuthService(store, cfg)
	authHandler := handler.NewAuthHandler(authService)

	// TODO: Setup Background Workers (The Collector)

	r := router.NewRouter(authHandler, cfg)

	log.Printf("Server running on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
