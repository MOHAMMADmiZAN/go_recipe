package api

import (
	"github.com/MOHAMMADmiZAN/go_recipe/internal/app/controller/ingredient"
	appMiddleware "github.com/MOHAMMADmiZAN/go_recipe/internal/app/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func IngredientApi() *mux.Router {
	ingredientApi := mux.NewRouter()
	ingredientApi.Handle("/ingredients", appMiddleware.AuthMiddleware(http.HandlerFunc(ingredient.CreateIngredient))).Methods(http.MethodPost)
	ingredientApi.HandleFunc("/ingredients", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello World"))

	}).Methods(http.MethodGet)

	return ingredientApi

}
