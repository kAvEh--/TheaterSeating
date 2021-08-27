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
