package cmd

import (
	"github.com/MOHAMMADmiZAN/go_recipe/internal/app/api"
	"github.com/MOHAMMADmiZAN/go_recipe/internal/pkg/utils"
	"os"
)

func StartApp() {
	utils.LoadEnv()
	config := api.APIConfig{
		Port:            os.Getenv("PORT"),
		SwaggerSpecPath: "./swagger.yaml",
	}
	api.RunAPIServer(config)

}
