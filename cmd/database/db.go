package database

import (
	"github.com/kAvEh--/TheaterSeating/cmd/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Initialize(dsn string) {
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Info(err)
		panic("db connection failed")
	}

	db.AutoMigrate(&model.User{}, &model.DBHall{}, &model.DBRow{}, &model.Seat{})
}

func CreateHall(hall model.Hall) {
	db.Create(hall)
}

func GetHall(id int) model.Hall {
	var hall model.DBHall
	db.Where("id = ?", id).First(&hall)
	ret := model.Hall{
		ID:   hall.ID,
		Name: hall.Name,
	}
	direction := false
	ranks := make(map[string][]model.Row)
	for i := 0; i < len(hall.Rows); i++ {
		lastRank := ""
		var row model.Row
		for j := 0; j < len(hall.Rows[i].Seats); j++ {
			if lastRank == "" || lastRank != hall.Rows[i].Seats[j].Rank {
				if len(row.Seats) > 0 {
					row.EmptySeatNum = len(row.Seats)
					if _, ok := ranks[hall.Rows[i].Seats[j].Rank]; ok {
						ranks[hall.Rows[i].Seats[j].Rank] = append(ranks[hall.Rows[i].Seats[j].Rank], row)
					}
				}
				lastRank = hall.Rows[i].Seats[j].Rank
				row = model.Row{
					ID:     hall.Rows[i].ID,
					IsRTL:  direction,
					Number: hall.Rows[i].Number,
					Seats:  make([]model.Seat, 0),
				}
			}

			if hall.Rows[i].Seats[j].State == 0 {
				row.Seats = append(row.Seats, hall.Rows[i].Seats[j])
			}
		}

		direction = !direction
	}
	return ret
}

func ReserveSeat(userID uint, seatID uint) {
	var seat model.Seat
	seat.ID = seatID
	db.First(&seat)
	seat.State = 1
	seat.User = int(userID)
	db.Save(&seat)
	var user model.User
	user.ID = userID
	db.First(&user)
	if user.Seats == nil {
		user.Seats = make([]model.Seat, 0)
	}
	user.Seats = append(user.Seats, seat)
	db.Save(&user)
}

func GetUser(userID uint) model.User {
	var user model.User
	user.ID = userID
	db.First(&user)

	return user
}
