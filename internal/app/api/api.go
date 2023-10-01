package api

import (
	"github.com/MOHAMMADmiZAN/go_recipe/internal/app/server"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/appResponse"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/db"
	"github.com/go-openapi/runtime/middleware"
	"log"
	"net/http"
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

// Config represents the API configuration.
type Config struct {
	Port        string
	SwaggerSpec string
}

// CreateAPIRouter creates a new API router.
func CreateAPIRouter(config Config) *http.ServeMux {
	// Create a new router
	router := http.NewServeMux()
	// Add routes
	router.HandleFunc("/health", HandleHealthRequest)
	router.Handle("/public/", ServeStaticFiles())

	// Swagger spec file
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	// Swagger UI docs
	swaggerOpts := middleware.SwaggerUIOpts{
		SpecURL: config.SwaggerSpec,
		Path:    "/docs",
	}
	swaggerMiddleware := middleware.SwaggerUI(swaggerOpts, nil)
	router.Handle("/docs", swaggerMiddleware)

	// Redoc
	redocOpts := middleware.RedocOpts{
		SpecURL: config.SwaggerSpec,
		Path:    "/redoc",
	}
	redocMiddleware := middleware.Redoc(redocOpts, nil)
	router.Handle("/redoc", redocMiddleware)

	return router
}

// RunAPIServer starts the API server.
func RunAPIServer(config Config) {
	router := CreateAPIRouter(config)
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
		server.RunServer(config.Port, router)
	}

}
