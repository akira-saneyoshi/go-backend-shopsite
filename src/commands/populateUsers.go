package main

import (
	"go-shopsite/src/database"
	"go-shopsite/src/models"

	"github.com/bxcodec/faker/v4"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		ambassador := models.User{
			FirstName:    faker.FirstName(),
			LastName:     faker.LastName(),
			Email:        faker.Email(),
			IsAmbassador: true,
		}

		ambassador.SetPassword("1234")

		database.DB.Create(&ambassador)
	}
}
