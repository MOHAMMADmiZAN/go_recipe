package api

import (
	"log"
	"net/http"
	"os"

	"github.com/MOHAMMADmiZAN/go_recipe/internal/app/server"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/appResponse"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/db"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/utils"
)

// HealthResponse represents a health response.
type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// HandleHealthRequest handles health endpoint requests.
func HandleHealthRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet || r.URL.Path != "/health" {
		http.NotFound(w, r)
		return
	}

	response := HealthResponse{
		Status:  "OK",
		Message: "The server is running",
	}

	appResponse.ResponseMessage(w, http.StatusOK, response)
}

// ServeStaticFiles Define a function named "ServeStaticFiles" to serve static files from the "/public/" path.
func ServeStaticFiles() http.Handler {
	return http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))
}

// RunAPIServer starts the API server.
func RunAPIServer() {
	utils.LoadEnv()
	port := os.Getenv("PORT")
	router := http.NewServeMux()

	router.HandleFunc("/health", HandleHealthRequest)
	router.Handle("/public/", ServeStaticFiles())

	err := db.Init()
	if err != nil {
		log.Printf("Error initializing database: %v", err)
		return
	}

	client, err := db.GetClient()
	if err != nil {
		log.Fatalf("Error getting database client: %v", err)
		return
	}
	if client != nil {
		server.RunServer(port, router)
	}

}
