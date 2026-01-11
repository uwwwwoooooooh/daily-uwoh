# Build Stage
FROM golang:1.25.5-alpine3.23 AS builder

WORKDIR /app
COPY . .

# Install git for fetching dependencies
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o daily-uwoh ./cmd/server
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.19.1/migrate.linux-amd64.tar.gz | tar xvz

# Run Stage
FROM alpine:3.23

WORKDIR /app

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/daily-uwoh .
COPY --from=builder /app/migrate .
COPY --from=builder /app/app.env . 
COPY --from=builder /app/start.sh .
COPY --from=builder /app/wait-for.sh .
COPY --from=builder /app/internal/db/migration ./internal/db/migration

RUN chmod +x /app/wait-for.sh /app/start.sh

EXPOSE 8080
CMD ["/app/daily-uwoh"]
ENTRYPOINT ["/app/start.sh"]
