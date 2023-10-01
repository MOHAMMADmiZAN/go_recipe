package api

import (
	"github.com/MOHAMMADmiZAN/go_recipe/internal/app/server"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/appResponse"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/db"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
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

// Config represents the API configuration.
type Config struct {
	Port        string
	SwaggerSpec string
}

// CreateAPIRouter creates a new API router.
func CreateAPIRouter(config Config) *mux.Router {
	// Create a new router
	router := mux.NewRouter()
	// Add routes
	router.HandleFunc("/health", HandleHealthRequest).Methods(http.MethodGet)
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// api v1 router group
	apiV1Router := router.PathPrefix("/api/v1").Subrouter()
	// auth router
	apiV1Router.PathPrefix("/auth").Handler(http.StripPrefix("/api/v1/auth", AuthApi()))

	// Swagger UI and Redoc setup
	swaggerOpts := middleware.SwaggerUIOpts{
		SpecURL: config.SwaggerSpec,
		Path:    "/docs",
	}
	router.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
	router.PathPrefix("/docs").Handler(middleware.SwaggerUI(swaggerOpts, nil))

	redocOpts := middleware.RedocOpts{
		SpecURL: config.SwaggerSpec,
		Path:    "/redoc",
	}
	router.PathPrefix("/redoc").Handler(middleware.Redoc(redocOpts, nil))
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
