package handlers

import (
	"encoding/json"

	"github.com/eduardofrnkdev/via-cep-golang/internal/services"
	"github.com/gorilla/mux"

	"net/http"
)

// SetupRoutes initializes the router and defines all routes.
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Health check route
	router.HandleFunc("/", HealthHandler)

	// CEP route
	router.HandleFunc("/cep/{id}", GetCepHandler)

	return router
}

// HealthHandler handle ping.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ping"))
}

// GetCepHandler handles requests to fetch CEP information.
func GetCepHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")
	cep, err := services.GetCep(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error(), "Error fetching CEP")
		return
	}
	w.Write([]byte(cep))
}

// respondWithError sends an error response with a JSON payload.
func respondWithError(w http.ResponseWriter, code int, errorMsg, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{
		"error":   errorMsg,
		"message": message,
	})
}
