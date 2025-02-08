package main

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go:embed gui/build/dist/*
var static embed.FS

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {
	// Strip the "gui/build/dist" prefix from the embedded files
	stripped, err := fs.Sub(static, "gui/build/dist")
	if err != nil {
		log.Fatal(err)
	}

	// Create a file server handler
	staticHandler := http.FileServer(http.FS(stripped))

	// Handle all routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check if file exists in embedded filesystem
		f, err := stripped.Open(r.URL.Path)
		if os.IsNotExist(err) {
			// Serve index.html for non-existent paths (SPA routing)
			r.URL.Path = "/"
		} else if err == nil {
			f.Close()
		}
		staticHandler.ServeHTTP(w, r)
	})

	apiServer := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	done := make(chan bool, 1)
	go gracefulShutdown(apiServer, done)

	log.Println("Server starting on :8080")
	if err := apiServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe(): %v", err)
	}

	<-done
	log.Println("Server stopped")
}
