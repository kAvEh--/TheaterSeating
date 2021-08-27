package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gofiber/fiber/v2"
	"github.com/kAvEh--/TheaterSeating/cmd/database"
	"github.com/kAvEh--/TheaterSeating/cmd/seating"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"strconv"
)

func main() {

	prefix := "server"
	viper.SetEnvPrefix(prefix)
	viper.AutomaticEnv()
	viper.SetConfigName(prefix + "_config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Debug("Config file changed:", e.Name)
	})

	dsn := viper.GetString("database_url")
	database.Initialize(dsn)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/user/:id", GetUser)
	app.Post("/reserve/:hall", ReserveSeat)
	app.Put("/hall", AddHall)
	app.Get("/hal/:id", GetHall)
	app.Listen(":3000")

	//h := seating.Hall{
	//	Ranks: map[string]seating.Rank{
	//		"red": {
	//			Rows: []seating.Row{
	//				{EmptySeatNum: 8,
	//					Seats: []database.Seat{
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//					},
	//				},
	//				{EmptySeatNum: 8,
	//					IsRTL: true,
	//					Seats: []database.Seat{
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//					},
	//				},
	//				{EmptySeatNum: 8,
	//					Seats: []database.Seat{
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//						{State: 0},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}

	//h = *seating.FindBestAlgorithm(h, "red", []int{1, 3, 4, 4, 5, 1, 4, 2})
	//for i := 0; i < len(h.Ranks["red"].Rows); i++ {
	//	for j := 0; j < len(h.Ranks["red"].Rows[i].Seats); j++ {
	//		fmt.Print(h.Ranks["red"].Rows[i].Seats[j].User, " ")
	//	}
	//	fmt.Println()
	//}
}

func GetHall(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.SendStatus(404)
	}
	h := database.GetHall(id)
	return c.JSON(h)
}

func AddHall(c *fiber.Ctx) error {
	h := new(database.DBHall)
	if err := c.BodyParser(h); err != nil {
		return err
	}

	logrus.Info("--->", len(h.Rows))
	logrus.Info()

	database.CreateHall(h)

	return c.SendStatus(200)
}

func GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.SendStatus(400)
	}
	user := database.GetUser(uint(id))
	return c.JSON(user)
}

func ReserveSeat(c *fiber.Ctx) error {
	h := c.Params("hall")
	id, err := strconv.Atoi(h)
	if err != nil {
		return c.SendStatus(400)
	}
	hall := database.GetHall(id)
	r := new(Reserve)

	if err := c.BodyParser(r); err != nil {
		return err
	}

	h1 := seating.FindBestAlgorithm(hall, r.Rank, r.Users)
	for _, v := range h1.Ranks {
		for i := 0; i < len(v); i++ {
			for j := 0; j < len(v[i].Seats); j++ {
				if v[i].Seats[j].State == 1 {
					database.ReserveSeat(uint(v[i].Seats[j].User), v[i].Seats[j].ID)
				}
			}
		}
	}

	return c.SendStatus(200)
}

type Reserve struct {
	Rank  string `json:"rank" xml:"rank" form:"rank"`
	Users []int  `json:"users" xml:"users" form:"users"`
}
