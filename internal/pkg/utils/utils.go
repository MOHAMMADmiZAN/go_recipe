package utils

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/appResponse"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LoadEnv load env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}
}

// LinkDefinition represents a link definition.
type LinkDefinition struct {
	Rel    string
	Href   string
	Method string
}

// LinkGenerator generates dynamic resource links based on a base URL and link definitions.
type LinkGenerator struct {
	BaseURL         string
	LinkDefinitions map[string]LinkDefinition
}

// NewLinkGenerator creates a new LinkGenerator instance.
func NewLinkGenerator(baseURL string, linkDefinitions map[string]LinkDefinition) *LinkGenerator {
	defaultDefinitions := map[string]LinkDefinition{
		"self": {
			Rel:    "self",
			Href:   "",
			Method: "GET",
		},
		"update": {
			Rel:    "update",
			Href:   "",
			Method: "PUT",
		},
		"delete": {
			Rel:    "delete",
			Href:   "",
			Method: "DELETE",
		},
	}

	// If linkDefinitions is nil, use defaultDefinitions
	if linkDefinitions == nil {
		linkDefinitions = defaultDefinitions
	}

	return &LinkGenerator{
		BaseURL:         baseURL,
		LinkDefinitions: linkDefinitions,
	}
}

// GenerateLinks generates dynamic resource links for a given resource ID.
func (lg *LinkGenerator) GenerateLinks(resourceID string) map[string]LinkDefinition {
	generatedLinks := make(map[string]LinkDefinition)

	for key, linkDefinition := range lg.LinkDefinitions {
		generatedLinks[key] = LinkDefinition{
			Rel:    linkDefinition.Rel,
			Href:   fmt.Sprintf("%s/%s%s", lg.BaseURL, resourceID, linkDefinition.Href),
			Method: linkDefinition.Method,
		}
	}

	return generatedLinks
}

// HexToObjectId hex to ObjectId
func HexToObjectId(hex string) primitive.ObjectID {
	var w http.ResponseWriter
	id, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		appResponse.ResponseMessage(w, http.StatusBadRequest, "ObjectId Create Failed")
	}
	return id
}
