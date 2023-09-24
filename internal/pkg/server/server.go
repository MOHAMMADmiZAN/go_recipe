package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/response"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RunServer(port string, handler http.Handler) {
	server := &http.Server{
		Addr:    ":" + port,
		Handler: response.ErrorHandler(handler),
	}

	go func() {
		fmt.Printf("Server is running on port %s...\n", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("Server stopped: %v\n", err)
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Error during shutdown: %v\n", err)
	}
	close(quit)
}
