package main

import (
	"log"

	"github.com/uwwwwoooooooh/daily-uwoh/internal/api/handler"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/api/router"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/db"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/repository"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/service"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/token"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/utils"
)

func main() {
	log.Println("Initializing DailyUwoh System...")

	cfg, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	log.Printf("Loaded config. Port: %s", cfg.ServerPort)

	if cfg.TokenSymmetricKey == "secret" {
		log.Println("WARNING: Using default TokenSymmetricKey 'secret'. This is insecure for production!")
	}

	db, err := db.ConnectDB(cfg.DBUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	tokenMaker, err := token.NewPasetoMaker(cfg.TokenSymmetricKey)
	if err != nil {
		log.Fatalf("cannot create token maker: %v", err)
	}

	store := repository.NewStore(db)
	authService := service.NewAuthService(store, tokenMaker, cfg)
	authHandler := handler.NewAuthHandler(authService)

	// TODO: Setup Background Workers (The Collector)

	r := router.NewRouter(authHandler, tokenMaker, cfg)

	r.SetTrustedProxies(nil)

	log.Printf("Server running on port %s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
