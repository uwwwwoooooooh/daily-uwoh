package main

import (
	"log"

	"github.com/uwwwwoooooooh/daily-uwoh/internal/config"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/database"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/router"
)

// =================================================================================
// ARCHITECTURE NOTES (Draft v0.1)
//
// System Design (Digital Editor):
// 1. Collector: Scrapes content (Twitter/Pixiv) -> Repository.
// 2. Processor: AI Analysis (Gemini/Vision) for tagging, scoring (Uwoh logic), and NSFW.
// 3. Storage: Postgres (JSONB Metadata) + File System/S3.
// 4. Publisher: Telegram Bot for curated content distribution.
// =================================================================================

// ---------------------------------------------------------------------------------
// 1. Domain Models (Data Structures)
// TODO: moved to 'model' package.
// ---------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------
// 2. Interfaces (Contracts)
// Define behavior before implementation. This makes unit testing easier (Mocking).
// ---------------------------------------------------------------------------------

// Moved to internal/repository and internal/service

// ---------------------------------------------------------------------------------
// 4. Main Application Entry
// ---------------------------------------------------------------------------------

func main() {
	log.Println("🛠️  Initializing DailyUwoh System...")

	// Step 1: Configuration
	cfg := config.LoadConfig()
	log.Printf("Loaded config. Port: %s", cfg.ServerPort)

	// Step 2: Database Connection
	db, err := database.ConnectDB(cfg.DBUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	_ = db // Keep compiler happy for now

	// Step 3: Migration
	// TODO: Run AutoMigrate for Artist, Artwork, Tag.
	log.Println("TODO: Run database migrations...")

	// Step 4: Setup Components (Dependency Injection)
	// repo := repository.NewPostgresArtworkRepository(db)
	// ai := processor.NewGeminiProcessor(apiKey)
	// pub := publisher.NewTelegramPublisher(botToken, channelID)

	// Step 5: Setup Background Workers (The Collector)
	// collector := service.NewCollector(repo, ai, pub)
	// go collector.Start(10) // Launch 10 workers

	// Step 6: HTTP Server (Gin)
	r := router.NewRouter()
	_ = r

	// Start server (blocking)
	// r.Run(":" + cfg.ServerPort)

	// Step 7: Graceful Shutdown
	// TODO: Listen for OS signals (SIGINT, SIGTERM) to close DB connections and stop workers safely.
	// This is critical for preventing data corruption during deployment.

	log.Println("🚧 Server not started. This is just a blueprint.")
	log.Println("Run `go run main.go` to verify compilation.")

	// Block forever for now, just to simulate a running process
	select {}
}

// ---------------------------------------------------------------------------------
// 5. Placeholders (To be implemented)
// ---------------------------------------------------------------------------------
