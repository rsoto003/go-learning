package main

import (
	"github.com/rsoto003/go-learning/initializers"
	"github.com/rsoto003/go-learning/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
