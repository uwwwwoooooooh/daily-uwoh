package service

// ScraperService defines the interface for the crawler.
// This will likely involve Goroutines and Channels.
type ScraperService interface {
	// Enqueue adds a URL to the scraping queue.
	Enqueue(url string)
	// Start launches the worker pool.
	Start(workerCount int)
}
