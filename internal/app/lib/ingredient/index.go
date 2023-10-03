package ingredientService

import (
	"github.com/MOHAMMADmiZAN/go_recipe/internal/app/model"
	"github.com/kamva/mgm/v3"
)

// create a new ingredient

func CreateIngredient(name string, description string, category string) (ingredient *model.Ingredient, err error) {
	ingredient = model.IngredientModel(name, description, category)
	err = mgm.Coll(ingredient).Create(ingredient)
	if err != nil {
		return nil, err
	}
	return ingredient, nil
}
