package main

import (
	"butter/app"
	"butter/helper"
	"butter/initializers"
	"butter/pkg/model/usermodel"
	"flag"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	handleArgs()
}

func handleArgs() {
	flag.Parse()
	db := app.NewDb()
	args := flag.Args()

	if len(args) >= 1 {
		switch args[0] {
		case "seed":
			var users []usermodel.UserEntity
			for i := 0; i < 100; i++ {
				username := "seed user " + strconv.Itoa(i)
				users = append(users, usermodel.UserEntity{
					ID:       uuid.New().String(),
					Username: username,
					Password: username,
					Email:    username,
					Name:     username,
				})
			}

			err := db.Create(&users).Error
			helper.PanicIfError(err)
		case "unseed":
			err := db.Unscoped().Delete(&usermodel.UserEntity{}, "email LIKE ?", "seed user%").Error
			helper.PanicIfError(err)
		}

	}
}

// How to Run:
// go run main.go seed
// go run main.go unseed
