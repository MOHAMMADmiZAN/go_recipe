package api

import (
	"encoding/json"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/server"
	"net/http"
)

// Define a function named "HandleHealthRequest" to handle health endpoint requests.
func HandleHealthRequest(w http.ResponseWriter, r *http.Request) {
	// If the HTTP method is GET and the path is "/health", return JSON data.
	if r.Method == http.MethodGet && r.URL.Path == "/health" {
		// Create a map to store the response data.
		response := map[string]interface{}{
			"status": "ok",
		}

		// Encode the map into JSON and send it as the response.
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// ServeStaticFiles Define a function named "ServeStaticFiles" to serve static files from the "/public/" path.
func ServeStaticFiles() http.Handler {
	return http.StripPrefix("/public/", http.FileServer(http.Dir("./static")))
}

// RunAPIServer Define a function named "RunAPIServer" to start the API server.
func RunAPIServer() {
	router := http.NewServeMux()

	// Handle the health endpoint.
	router.HandleFunc("/health", HandleHealthRequest)

	// Serve static files from the "/public/" path.
	router.Handle("/public/", ServeStaticFiles())

	// Start the server on port 8080.
	server.RunServer("8080", router)
}
