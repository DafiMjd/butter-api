package main

import (
	"butter/initializers"
	"os"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	app := InitializedServer()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	err := app.Listen("0.0.0.0:" + port)
	if err != nil {
		panic(err)
	}
}
