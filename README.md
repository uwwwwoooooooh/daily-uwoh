# DailyUwoh 😭😭😭😭😭😭😭

**Goal: An automated "Digital Editor" that hunts, curates, and archives anime illustrations.**
This system acts as an autonomous editor that:
1.  **Collects**: Scrapes Twitter/X for new art from tracked artists.
2.  **Reviews**: Uses AI (Vision) to judge if the art is "Uwoh" enough (Cute/Sexy) and flags NSFW content.
3.  **Publishes**: Pushes the best filtered content to a Telegram Channel.
4.  **Archives**: Saves everything to specific deduplicated storage.

**Data Flow:**
`Twitter (Source) -> Go Backend (Collector) -> AI Vision (Processor) -> Database (Storage) -> Telegram Bot (Publisher)`

* **`github.com/gin-gonic/gin`**: Web framework.
* **`github.com/jackc/pgx/v5`**: PostgreSQL driver and toolkit.
* **`github.com/sqlc-dev/sqlc`**: Type-safe SQL compiler for Go.
* **`net/http` & Context**: For controlling high-concurrency scraping.


```text
DailyUwoh/
├── cmd/
│   └── server/          # Application entry point (main.go)
├── internal/            # Private application logic
│   ├── api/             # API Layer (Handler, Router, Middleware)
│   ├── config/          # Configuration management (.env loading)
│   ├── db/              # Database (Connection, Migration, Queries, SQLC)
│   ├── model/           # Domain Models
│   ├── processor/       # AI Image Analysis (Gemini/OpenAI)
│   ├── publisher/       # Content Distribution (Telegram)
│   ├── repository/      # Data Access Layer (Repository Pattern)
│   └── service/         # Core Business Logic
├── deployments/         # Docker & CI/CD configurations
│   └── Dockerfile       # Docker build configuration
├── .gitignore           # Git ignore rules
├── docker-compose.yml   # Docker services setup
├── go.mod               # Go module dependencies
└── README.md            # Documentation
```
