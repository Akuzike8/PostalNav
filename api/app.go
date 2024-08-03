package main

import (
	"api/handles"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Serve static files
	fs := http.FileServer(http.Dir("../static"))
	r.Handle("/doc/*", http.StripPrefix("/doc/", fs))

	r.Get("/", handles.Home)

	r.Get("/get", handles.GetPostal)

	r.Get("/get/{slug:[a-z-0-9-]+}", handles.GetPostalBySlug)

	error := godotenv.Load(".env")

	if error != nil {
		fmt.Println(error)
		log.Fatal("Error loading .env file")
	}

	http.ListenAndServe(":8080", r)
}
