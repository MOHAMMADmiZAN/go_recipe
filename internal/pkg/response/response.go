package response

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

// JsonResponse Define a helper function to send JSON responses
func JsonResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	return enc.Encode(data)
}

// HandleError Helper function to handle errors
func HandleError(w http.ResponseWriter, r *http.Request, err interface{}) {
	// Get the stack trace
	stack := make([]byte, 1024)
	runtime.Stack(stack, false)

	// Log the error and the stack trace
	logger := log.New(os.Stderr, "ERROR: ", log.LstdFlags)
	logger.Printf("Recovered from panic: %v\n%s", err, stack)

	// Send a 500 Internal Server Error response
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
}

// HandleErrorAndExit Helper function to handle errors and exit the application
func HandleErrorAndExit(message string, err error, exitCode int) {
	fmt.Printf("ERROR: %s: %v\n", message, err)
	os.Exit(exitCode)
}

func ErrorHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				HandleError(w, r, err)
			}
		}()
		handler.ServeHTTP(w, r)
	})
}
