package main

import (
	"context"
	"log"

	"github.com/uwwwwoooooooh/daily-uwoh/internal/model"
)

// =================================================================================
// ARCHITECTURE NOTES (Draft v0.1)
//
// System Design:
// 1. Layered Architecture: Handler (HTTP) -> Service (Business Logic) -> Repository (Data Access).
// 2. Concurrency: Use a Worker Pool pattern for the scraper (Producer-Consumer).
// 3. Database: Postgres with GORM. Heavy use of JSONB for dynamic metadata.
// 4. External: C++ Microservice via gRPC for pHash calculation (Future).
// =================================================================================

// ---------------------------------------------------------------------------------
// 1. Domain Models (Data Structures)
// TODO: moved to 'model' package.
// ---------------------------------------------------------------------------------

// ---------------------------------------------------------------------------------
// 2. Interfaces (Contracts)
// Define behavior before implementation. This makes unit testing easier (Mocking).
// ---------------------------------------------------------------------------------

// ArtworkRepository defines how we interact with the database.
type ArtworkRepository interface {
	Create(ctx context.Context, artwork *model.Artwork) error
	FindByHash(ctx context.Context, hash string) (*model.Artwork, error)
	// TODO: Add search with pagination
}

// ScraperService defines the interface for the crawler.
// This will likely involve Goroutines and Channels.
type ScraperService interface {
	// Enqueue adds a URL to the scraping queue.
	Enqueue(url string)
	// Start launches the worker pool.
	Start(workerCount int)
}

// ---------------------------------------------------------------------------------
// 3. Configuration & Global State
// ---------------------------------------------------------------------------------

type Config struct {
	DBUrl      string
	ServerPort string
	// TODO: Add Redis config for job queue
	// TODO: Add gRPC address for C++ service
}

// LoadConfig reads from .env or flags.
// Note: For now, hardcode values for quick testing.
func LoadConfig() Config {
	return Config{
		DBUrl:      "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable",
		ServerPort: ":8080",
	}
}

// ---------------------------------------------------------------------------------
// 4. Main Application Entry
// ---------------------------------------------------------------------------------

func main() {
	log.Println("🛠️  Initializing DailyUwoh System...")

	// Step 1: Configuration
	cfg := LoadConfig()
	log.Printf("Loaded config. Port: %s", cfg.ServerPort)

	// Step 2: Database Connection
	// TODO: Implement GORM connection logic here.
	// Note: Remember to set SetMaxOpenConns and SetMaxIdleConns for connection pooling.
	// db := connectDB(cfg.DBUrl)

	// Step 3: Migration
	// TODO: Run AutoMigrate for Artist, Artwork, Tag.
	log.Println("TODO: Run database migrations...")

	// Step 4: Setup Components (Dependency Injection)
	// repo := NewPostgresArtworkRepository(db)
	// service := NewCoreService(repo)

	// Step 5: Setup Background Workers (The Crawler)
	// scraper := NewScraperService()
	// go scraper.Start(10) // Launch 10 workers in background

	// Step 6: HTTP Server (Gin)
	// r := gin.Default()
	// setupRoutes(r)

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

func setupRoutes(r interface{}) {
	// TODO: Define API groups
	// v1 := r.Group("/api/v1")
	// v1.POST("/upload", handleUpload)
	// v1.GET("/feed", handleFeed)
}
