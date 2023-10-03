package ingredient

import (
	"encoding/json"
	ingredientService "github.com/MOHAMMADmiZAN/go_recipe/internal/app/lib/ingredient"
	appMiddleware "github.com/MOHAMMADmiZAN/go_recipe/internal/app/middleware"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/appResponse"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/validator"
	"net/http"
)

type createRequestIngredient struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Category    string `json:"category" validate:"required"`
}

func CreateIngredient(w http.ResponseWriter, r *http.Request) {
	// Get the user information from the request context

	_, ok := r.Context().Value("user").(*appMiddleware.User)
	if !ok {
		appResponse.ResponseMessage(w, http.StatusUnauthorized, "Unauthorized")
	}

	// decode the request body into a struct
	var newIngredient createRequestIngredient
	if err := json.NewDecoder(r.Body).Decode(&newIngredient); err != nil {
		appResponse.ResponseMessage(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	// validate the request body
	if err := validator.ValidateStruct(newIngredient); err != nil {
		appResponse.ResponseMessage(w, http.StatusBadRequest, err)
		return
	}
	// create ingredient
	ingredient, err := ingredientService.CreateIngredient(newIngredient.Name, newIngredient.Description, newIngredient.Category)
	if err != nil {
		appResponse.ResponseMessage(w, http.StatusInternalServerError, err)
		return
	}
	// return the created ingredient
	appResponse.ResponseMessage(w, http.StatusCreated, ingredient)

}
