package main

import (
	"go-shopsite/src/database"
	"go-shopsite/src/models"
	"math/rand"

	"github.com/bxcodec/faker/v4"
)

func main() {
	database.Connect()

	for i := 0; i < 30; i++ {
		var OrderItems []models.OrderItem

		for j := 0; j < rand.Intn(5); j++ {
			price := float64(rand.Intn(90) + 10)
			qty := uint(rand.Intn(5))

			OrderItems = append(OrderItems, models.OrderItem{
				ProductTitle:      faker.Word(),
				Price:             price,
				Quantity:          qty,
				AdminRevenue:      0.9 * price * float64(qty),
				AmbassadorRevenue: 0.1 * price * float64(qty),
			})
		}

		database.DB.Create(&models.Order{
			UserId:          uint(rand.Intn(30) + 1),
			Code:            faker.Username(),
			AmbassadorEmail: faker.Email(),
			FirstName:       faker.FirstName(),
			LastName:        faker.LastName(),
			Email:           faker.Email(),
			Complete:        true,
			OrderItems:      OrderItems,
		})
	}
}
