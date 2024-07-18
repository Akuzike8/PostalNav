package main

import (
	"net/http"
	"api/handles"
	"fmt"
	"log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

    // Serve static files
	fs := http.FileServer(http.Dir("../static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", handles.Home)

    r.Post("/add", handles.AddPostal)

	r.Get("/get", handles.GetPostal)

	error := godotenv.Load(".env")

    if error != nil {
        fmt.Println(error)
        log.Fatal("Error loading .env file")
    }

	http.ListenAndServe(":8080", r)
}
