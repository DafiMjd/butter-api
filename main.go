package main

import (
	"butter/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

// func main() {
// 	app := InitializedServer()

// 	port := os.Getenv("PORT")

// 	if port == "" {
// 		port = "3000"
// 	}

//		err := app.Listen("0.0.0.0:" + port)
//		if err != nil {
//			panic(err)
//		}
//	}
func main() {
	app := InitializedServer()

	err := app.Listen("localhost:5000")
	if err != nil {
		panic(err)
	}
}
