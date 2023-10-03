package model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	mgm.DefaultModel `bson:",inline"`
	Title            string             `json:"title" bson:"title"`
	Ingredients      []RecipeIngredient `json:"ingredients" bson:"ingredients"`
	Instructions     string             `json:"instructions" bson:"instructions"`
	Creator          primitive.ObjectID `json:"creator" bson:"creator"`
}

type RecipeIngredient struct {
	IngredientID primitive.ObjectID `json:"ingredient" bson:"ingredient"`
	Quantity     string             `json:"quantity" bson:"quantity"`
}

func RecipeModel(title string, ingredients []RecipeIngredient, instructions string, creator primitive.ObjectID) *Recipe {
	return &Recipe{
		Title:        title,
		Ingredients:  ingredients,
		Instructions: instructions,
		Creator:      creator,
	}
}

func (r *Recipe) CollectionName() string {
	return "recipes"
}
