package api

import (
	"github.com/MOHAMMADmiZAN/go_recipe/internal/app/controller/auth"
	"github.com/gorilla/mux"
	"net/http"
)

func AuthApi() *mux.Router {
	authApi := mux.NewRouter()
	authApi.HandleFunc("/signup", auth.SignUP).Methods(http.MethodPost)
	authApi.HandleFunc("/signin", auth.SignIN).Methods(http.MethodPost)

	return authApi
}
