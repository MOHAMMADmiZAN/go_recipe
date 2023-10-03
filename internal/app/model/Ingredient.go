package model

import "github.com/kamva/mgm/v3"

type Ingredient struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Description      string `json:"description" bson:"description"`
	Category         string `json:"category" bson:"category"`
}

func IngredientModel(name string, description string, category string) *Ingredient {
	return &Ingredient{
		Name:        name,
		Description: description,
		Category:    category,
	}
}

func (i *Ingredient) CollectionName() string {
	return "ingredients"
}
