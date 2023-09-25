package utils

import (
	"github.com/joho/godotenv"
	"os"
)

// LoadEnv load env file
func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		os.Exit(1)
	}
}
