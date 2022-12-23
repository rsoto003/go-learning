package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

// func starting with capital is important if wanting to use func in other files.
func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
