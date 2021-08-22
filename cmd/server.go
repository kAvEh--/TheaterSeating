package main

import (
	"fmt"
	"github.com/gofiber/fiber/cmd/seating"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	//app.Listen(":3000")

	h := seating.Hall{
		Ranks: map[string]seating.Rank{
			"red": {
				Rows: []seating.Row{
					{Seats: []seating.Seat{
						{State: 0},
						{State: 0},
						{State: 0},
						{State: 0},
						{State: 0},
						{State: 0},
						{State: 0},
						{State: 0},
					},
					},
					{Seats: []seating.Seat{
						{State: 0},
						{State: 0},
						{State: 0},
						{State: 0},
						{State: 0},
						{State: 0},
						{State: 0},
						{State: 0},
					},
					},
					{Seats: []seating.Seat{
						{State: 0},
						{State: 0},
						{State: 0},
						{State: 0},
						{State: 0},
						{State: 0},
						{State: 0},
						{State: 0},
					},
					},
				},
			},
		},
	}

	h.FillSeating("red", []int{1, 3, 4, 4, 5, 1, 2, 4})
	for i := 0; i < len(h.Ranks["red"].Rows); i++ {
		for j := 0; j < len(h.Ranks["red"].Rows[i].Seats); j++ {
			fmt.Print(h.Ranks["red"].Rows[i].Seats[j].User, " ")
		}
		fmt.Println()
	}
}
