package server

import (
	"context"
	"errors"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/appError"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/db"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RunServer(port string, handler http.Handler) {
	server := &http.Server{
		Addr:    ":" + port,
		Handler: appError.ErrorHandler(handler),
	}

	go func() {
		log.Printf("Server is running on port %s...\n", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Server stopped: %v\n", err)
		}
	}()
	quit := make(chan struct{})
	gracefulShutdown(server, quit)
	<-quit
}

func gracefulShutdown(server *http.Server, quit chan struct{}) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer func() {
		err := db.Disconnect()
		if err != nil {
			log.Printf("Error disconnecting from database: %v\n", err)
			cancel()
		}
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error during shutdown: %v\n", err)
	}
	close(quit)
}
