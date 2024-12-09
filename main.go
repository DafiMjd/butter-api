package main

import (
	"os"
)

// func init() {
// 	initializers.LoadEnvVariables()
// }

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

// docker
// func main() {
// 	app := InitializedServer()

// 	err := app.Listen(":8080")
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func main() {
// 	http.HandleFunc("/", handlerPort)
// 	http.HandleFunc("/port", handlerPort)
// 	http.ListenAndServe(":8080", nil)
// }

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hellow World!!!!!\n")
// }

// func handlerPort(w http.ResponseWriter, r *http.Request) {
// 	port := os.Getenv("DB_PORT")
// 	user := os.Getenv("POSTGRES_USER")
// 	fmt.Fprintf(w, "Hellow World!!!!"+port+user)
// }
