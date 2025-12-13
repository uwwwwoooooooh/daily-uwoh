package main

import (
	"context"
	"log"

	"gorm.io/gorm"
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
// TODO: Eventually move these to a separate 'models' package to avoid circular deps.
// ---------------------------------------------------------------------------------

// Artist represents a creator on platforms like Twitter/Pixiv.
type Artist struct {
	gorm.Model
	Name string `json:"name"`

	// SocialProfiles (JSONB): Flexible storage for platform links.
	// Structure: {"twitter": "url", "pixiv": "url", "bluesky": "url"}
	// Note: We use interface{} here, but in production, we might want a specific struct for type safety.
	SocialProfiles map[string]interface{} `gorm:"type:jsonb;serializer:json"`

	Artworks []Artwork // Has-Many relationship
}

// Artwork represents the core entity of the system.
type Artwork struct {
	gorm.Model
	Title     string
	FilePath  string // Local storage path or S3 Key
	SourceURL string // Where did we scrape this from?

	// PHash: The "Perceptual Hash" calculated by our C++ module.
	// Important: We need a B-Tree index on this for Hamming Distance calculation later?
	// Or just a standard index for exact string matching for now.
	PHash string `gorm:"index"`

	// MetaData (JSONB): Raw data from the source (likes, retweets, dimensions).
	MetaData map[string]interface{} `gorm:"type:jsonb;serializer:json"`

	ArtistID uint
	Tags     []*Tag `gorm:"many2many:artwork_tags;"`
}

// Tag for categorization.
// TODO: Determine if we need a 'Confidence' field in the join table for AI confidence levels.
type Tag struct {
	gorm.Model
	Name     string `gorm:"uniqueIndex"`
	Category string // e.g., "character", "series", "artist_style"
}

// ---------------------------------------------------------------------------------
// 2. Interfaces (Contracts)
// Define behavior before implementation. This makes unit testing easier (Mocking).
// ---------------------------------------------------------------------------------

// ArtworkRepository defines how we interact with the database.
type ArtworkRepository interface {
	Create(ctx context.Context, artwork *Artwork) error
	FindByHash(ctx context.Context, hash string) (*Artwork, error)
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
