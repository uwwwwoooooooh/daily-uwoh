# DailyUwoh 😭😭😭😭😭😭😭

**Goal: An automated "Digital Editor" that hunts, curates, and archives anime illustrations.**

This system acts as an autonomous editor that:
1.  **Collects**: Scrapes Twitter for new art from tracked artists.
2.  **Reviews**: Uses AI (Vision) to judge if the art is "Uwoh" enough (Cute/Sexy) and flags NSFW content.
3.  **Publishes**: Pushes the best filtered content to a Telegram Channel.
4.  **Archives**: Saves everything to specific deduplicated storage.

**Data Flow:**
`Twitter (Source) -> Go Backend (Collector) -> AI Vision (Processor) -> Database (Storage) -> Telegram Bot (Publisher)`

## Tech Stack

*   **Language**: Go 1.23+
*   **Web Framework**: `github.com/gin-gonic/gin`
*   **Database**: PostgreSQL 16+ using `pgx/v5` driver.
*   **Data Access**: `github.com/sqlc-dev/sqlc` for type-safe SQL generation.
*   **Mocking**: `go.uber.org/mock/gomock` for unit testing.
*   **Config**: `github.com/spf13/viper` (via `utils` package).

## Project Structure

```text
DailyUwoh/
├── cmd/
│   └── server/          # Application entry point (main.go)
├── internal/            # Private application logic
│   ├── api/             # API Layer (Handler, Router, Middleware)
│   ├── db/              # Database (Connection, Migration, Queries, SQLC, Mock)
│   ├── model/           # Domain Models
│   ├── processor/       # AI Image Analysis (Gemini/OpenAI)
│   ├── publisher/       # Content Distribution (Telegram)
│   ├── repository/      # Data Access Layer (Repository Pattern)
│   ├── service/         # Core Business Logic
│   └── utils/           # Utilities & Config (.env loading, error handling)
├── deployments/         # Docker & CI/CD configurations
│   └── Dockerfile       # Docker build configuration
├── .gitignore           # Git ignore rules
├── app.env              # Environment variables (Example)
├── docker-compose.yml   # Docker services setup
├── go.mod               # Go module dependencies
└── Makefile             # Development commands
```

## Getting Started

### Prerequisites

*   **Go** 1.23 or higher
*   **Docker**
*   **Make**
*   **Migrate CLI** (for DB migrations)
    *   `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`
*   **Sqlc**
    *   `go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest`
*   **Mockgen**
    *   `go install go.uber.org/mock/mockgen@latest`

### Setup

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/uwwwwoooooooh/daily-uwoh.git
    cd daily-uwoh
    ```

2.  **Start PostgreSQL DB:**
    ```bash
    make postgres
    # Or manually via docker-compose
    docker-compose up -d postgres
    ```

3.  **Create Database:**
    ```bash
    make createdb
    ```

4.  **Run Migrations:**
    ```bash
    make migrateup
    ```

5.  **Run the Server:**
    ```bash
    make server
    ```

## Development

### Makefile Commands

| Command         | Description                                      |
| :-------------- | :----------------------------------------------- |
| `make build`    | Build the binary.                                |
| `make run`      | Run the application.                             |
| `make server`   | Run the application (alias).                     |
| `make test`     | Run all tests.                                   |
| `make clean`    | Clean build artifacts.                           |
| `make sqlc`     | Generate Go code from SQL queries.               |
| `make mock`     | Generate MockDB store for testing.               |
| `make createdb` | Create the PostgreSQL database (via Docker).     |
| `make dropdb`   | Drop the PostgreSQL database.                    |
| `make migrateup`| Apply database migrations.                       |
| `make migratedown`| Revert database migrations.                    |

### Testing

The project uses `gomock` for mocking database interactions in the API layer.

1.  **Generate Mocks** (if you change `Store` interface):
    ```bash
    make mock
    ```

2.  **Run Tests:**
    ```bash
    make test
    ```
