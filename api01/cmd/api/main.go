package main

import (
    "context"
    "io"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"

    "goapi/internal/api/repository/DAL/SQLite"
    "goapi/internal/api/server"
    "goapi/internal/api/service"
)

func NewSimpleLogger(logFile string) *log.Logger {
    file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        // fallback to stdout if file can't be opened
        return log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
    }
    return log.New(io.MultiWriter(file, os.Stdout), "", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
    logger := NewSimpleLogger("app.log")

    // Context with cancellation on OS signals
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    // Open or create SQLite DB and ensure schema exists
    db, err := SQLite.Open("production.db")
    if err != nil {
        logger.Fatalf("failed to open db: %v", err)
    }
    defer db.Close()

    if err := SQLite.EnsureSchema(db); err != nil {
        logger.Fatalf("failed to ensure schema: %v", err)
    }

    // Create service factory
    sf := service.NewServiceFactory(db, logger, ctx)

    // Initialize server
    srv := server.NewServer(ctx, sf, logger)

    // Graceful shutdown
    go func() {
        ch := make(chan os.Signal, 1)
        signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
        <-ch
        logger.Println("shutting down server...")
        _ = srv.Shutdown()
        cancel()
    }()

    addr := ":8080"
    logger.Printf("Starting server on %s", addr)
    if err := srv.ListenAndServe(addr); err != nil && err != http.ErrServerClosed {
        logger.Fatalf("server error: %v", err)
    }
}
