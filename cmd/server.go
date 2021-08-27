package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/gofiber/fiber/v2"
	"github.com/kAvEh--/TheaterSeating/cmd/database"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Listen(":3000")

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
	logrus.Info("------------->>>>>", dsn)
	database.Initialize(dsn)

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
