# DailyUwoh 😭😭😭😭😭😭😭

**Goal: To track, find, and secure my favorite Waifus across the internet.**

I need an automated hunter-seeker system that monitors artists and indexes everything instantly. If a new illustration drops, I want it found and cataloged before I even wake up.

```text
DailyUwoh/
├── cmd/
│   └── server/          # Application entry point (main.go)
├── internal/            # Private application logic
│   ├── model/           # Data entities & DB Schema (Structs)
│   ├── repository/      # Data Access Layer (Postgres CRUD operations)
│   ├── service/         # Business Logic (Scraping rules, C++ integration)
│   ├── handler/         # HTTP Controllers (Gin route handlers)
│   └── config/          # Configuration management (.env loading)
├── deployments/         # Docker & CI/CD configurations
├── go.mod               # Go module dependencies
└── README.md            # Documentation

* **`github.com/gin-gonic/gin`**: Web framework.
* **`gorm.io/gorm`**: ORM. It maps my complex obsession (Artists, Tags, Metadata) into **PostgreSQL** without me needing to write raw SQL.
* **`gorm.io/driver/postgres`**: Driver to talk to the DB.
* **`net/http` & `context`**

