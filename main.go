package main

import (
	"butter/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	app := InitializedServer()

	err := app.Listen("localhost:5000")
	if err != nil {
		panic(err)
	}
}
