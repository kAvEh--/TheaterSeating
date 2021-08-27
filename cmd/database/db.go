package database

import (
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

	err = db.AutoMigrate(&User{}, &DBHall{}, &DBRow{}, &Seat{})
	if err != nil {
		panic(err)
	}
}

func CreateHall(hall *DBHall) {
	db.Create(&hall)
}

func GetHall(id int) Hall {
	var hall DBHall
	db.Where("id = ?", id).First(&hall)
	ret := Hall{
		ID:   hall.ID,
		Name: hall.Name,
	}
	direction := false
	ranks := make(map[string][]Row)
	for i := 0; i < len(hall.Rows); i++ {
		lastRank := ""
		var row Row
		for j := 0; j < len(hall.Rows[i].Seats); j++ {
			if lastRank == "" || lastRank != hall.Rows[i].Seats[j].Rank {
				if len(row.Seats) > 0 {
					row.EmptySeatNum = len(row.Seats)
					if _, ok := ranks[hall.Rows[i].Seats[j].Rank]; ok {
						ranks[hall.Rows[i].Seats[j].Rank] = append(ranks[hall.Rows[i].Seats[j].Rank], row)
					} else {
						ranks[hall.Rows[i].Seats[j].Rank] = []Row{row}
					}
				}
				lastRank = hall.Rows[i].Seats[j].Rank
				row = Row{
					IsRTL:  direction,
					Number: hall.Rows[i].Number,
					Seats:  make([]Seat, 0),
				}
			}

			if hall.Rows[i].Seats[j].State == 0 {
				row.Seats = append(row.Seats, hall.Rows[i].Seats[j])
			}
		}

		direction = !direction
	}
	ret.Ranks = ranks
	return ret
}

func ReserveSeat(userID uint, seatID uint) {
	var seat Seat
	seat.ID = seatID
	db.First(&seat)
	seat.State = 1
	seat.User = int(userID)
	db.Save(&seat)
	var user User
	user.ID = userID
	db.First(&user)
	if user.Seats == nil {
		user.Seats = make([]Seat, 0)
	}
	user.Seats = append(user.Seats, seat)
	db.Save(&user)
}

func GetUser(userID uint) User {
	var user User
	user.ID = userID
	db.First(&user)

	return user
}
