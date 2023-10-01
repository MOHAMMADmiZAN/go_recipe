package cmd

import (
	"github.com/MOHAMMADmiZAN/go_recipe/internal/app/api"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/utils"
	"os"
)

func StartApp() {
	utils.LoadEnv()
	config := api.Config{
		Port:        os.Getenv("PORT"),
		SwaggerSpec: "swagger.yaml",
	}
	api.RunAPIServer(config)

}
