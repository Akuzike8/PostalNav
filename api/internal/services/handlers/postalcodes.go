package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func SearchPostalBySlug(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	if slug == "" {
		http.Error(w, "Slug is required", http.StatusBadRequest)
		return
	}

	postal, err := services.GetPostalBySlug(slug)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching postal code: %v", err), http.StatusInternalServerError)
		return
	}

	if postal == nil {
		http.Error(w, "Postal code not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(postal)
}

func GetPostal(w http.ResponseWriter, r *http.Request) {
	postalCode := r.URL.Query().Get("code")
	if postalCode == "" {
		http.Error(w, "Postal code is required", http.StatusBadRequest)
		return
	}

	postal, err := services.GetPostalByCode(postalCode)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching postal code: %v", err), http.StatusInternalServerError)
		return
	}

	if postal == nil {
		http.Error(w, "Postal code not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(postal)
}
