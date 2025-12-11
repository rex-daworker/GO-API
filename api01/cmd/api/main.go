// cmd/api/main.go (unchanged structure, only imports and factory changed)
package main

import (
	"context"
	"io" // <-- add this
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
		log.Fatal(err)
	}
	return log.New(io.MultiWriter(file, os.Stdout), "", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := NewSimpleLogger("production.log")
	db, err := SQLite.NewSqlite("production.db")
	if err != nil {
		logger.Println("Error setting up database:", err)
		return
	}
	defer db.Close()

	sf := service.NewServiceFactory(db, logger, ctx)
	srv := server.NewServer(ctx, sf, logger)

	gracefulShutdown(srv, cancel, logger)

	logger.Println("Starting server on :8080...")
	if err := srv.ListenAndServe(":8080"); err != nil {
		if err != http.ErrServerClosed {
			logger.Println("Server startup error:", err)
		}
		logger.Println("Server gracefully shutdown complete.")
	}
}

func gracefulShutdown(server *server.Server, cancel context.CancelFunc, logger *log.Logger) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-signalCh
		cancel()
		if err := server.Shutdown(); err != nil {
			logger.Println("Error shutting down API Server:", err)
		}
	}()
}
