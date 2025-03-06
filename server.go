package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"main/api/db"
	"main/api/routes"
)

func Start(ctx context.Context) error {
	if err := db.Initialize(); err != nil {
		return fmt.Errorf("failed to initialize database: %v", err)
	}
	defer db.Close()

	router := routes.SetupRoutes()

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// Channel to receive any errors returned by the server
	serverErrors := make(chan error, 1)

	// Start the server
	go func() {
		fmt.Printf("Starting server on port %s\n", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	// Wait for either context cancellation or server error
	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %v", err)
	case <-ctx.Done():
		// Shutdown the server
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		fmt.Println("Shutting down server...")
		if err := server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("server shutdown error: %v", err)
		}
		fmt.Println("Server gracefully stopped")
	}

	return nil
}
