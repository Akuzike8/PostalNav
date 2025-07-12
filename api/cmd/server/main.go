package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

    r := chi.NewRouter()
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/"))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

    // Authentication routes (public)
	r.Route("/auth", func(r chi.Router) {

	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%v", "8080"),
		Handler: r,
	}

	ch := make(chan error, 1)

	log.Println("Starting server on :8080")

	go func ()  {
		if err := server.ListenAndServe(); err != nil {
			ch <- fmt.Errorf("failed to start server:%w", err)
		}

		close(ch)
	}()

	select {
		case err := <-ch:
			log.Fatal(err.Error())
		case <-ctx.Done():
			timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			server.Shutdown(timeout)

	}
}
