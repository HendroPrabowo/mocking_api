package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mocking_api/dukcapil"
	"mocking_api/health"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/httplog"
	log "github.com/sirupsen/logrus"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := registerRoutes()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Create a channel to receive signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a separate goroutine
	go func() {
		log.Printf("server listening on port %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for a signal to shutdown the server
	sig := <-signalCh
	log.Println("received signal: %v", sig)

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server shutdown failed: %v\n", err)
	}

	log.Println("server shutdown gracefully")
}

func registerRoutes() *chi.Mux {
	r := chi.NewRouter()
	r = setCors(r)

	logger := httplog.NewLogger("mocking_api", httplog.Options{
		JSON: true,
	})
	r.Use(httplog.RequestLogger(logger))

	// REGISTER ALL ROUTES HERE
	// health check routes
	health.RegisterRoutes(r)

	// dukcapil routes
	dukcapilRoutes, _ := dukcapil.InitializeDukcapil()
	dukcapilRoutes.RegisterRoutes(r)

	return r
}

func setCors(r *chi.Mux) *chi.Mux {
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	cors := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
	return r
}

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}
