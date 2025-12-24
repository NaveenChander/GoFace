package main

import (
	"fmt"
	"os"

	"github.com/NaveenChander/GoFace/simulator/api"
	"github.com/NaveenChander/GoFace/simulator/db"
	"github.com/NaveenChander/GoFace/simulator/models"
)

func main() {
	fmt.Println("Hello, World!")

	modelsEnv := models.EnvConfig{}
	err := modelsEnv.LoadConfig()
	db.GetDBContext()

	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
		return
	}

	fmt.Printf("Loaded configuration: %+v\n", models.EnvironmentConfig)

	api.StartServer()

}
