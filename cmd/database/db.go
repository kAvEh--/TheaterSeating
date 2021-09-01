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
	h := DBHall{
		Name: hall.Name,
	}
	db.Create(&h)
	for i := 0; i < len(hall.DbRows); i++ {
		r := DBRow{
			Number:   hall.DbRows[i].Number,
			DBHallID: h.ID,
		}
		db.Create(&r)
		for j := 0; j < len(hall.DbRows[i].Seats); j++ {
			s := Seat{
				Rank:    hall.DbRows[i].Seats[j].Rank,
				DBRowID: r.ID,
				UserID:  0,
			}
			db.Create(&s)
			hall.DbRows[i].Seats[j] = s
		}
		hall.DbRows[i] = r
	}
	db.Model(&h).Association("DbRows").Replace(hall.DbRows)
}

func GetHall(id int) Hall {
	var hall DBHall
	db.Where("id = ?", id).First(&hall)
	var v []DBRow
	db.Where("db_hall_id = ?", hall.ID).Find(&v)
	hall.DbRows = v
	ret := Hall{
		ID:   hall.ID,
		Name: hall.Name,
	}
	direction := false
	ranks := make(map[string][]Row)
	for i := 0; i < len(hall.DbRows); i++ {
		lastRank := ""
		var row Row
		var seats []Seat
		db.Where("db_row_id = ?", hall.DbRows[i].ID).Find(&seats)
		hall.DbRows[i].Seats = seats
		for j := 0; j < len(hall.DbRows[i].Seats); j++ {
			if lastRank == "" || lastRank != hall.DbRows[i].Seats[j].Rank || j == len(hall.DbRows[i].Seats)-1 {
				if len(row.Seats) > 0 {
					row.EmptySeatNum = len(row.Seats)
					if _, ok := ranks[hall.DbRows[i].Seats[j].Rank]; ok {
						ranks[hall.DbRows[i].Seats[j].Rank] = append(ranks[hall.DbRows[i].Seats[j].Rank], row)
					} else {
						ranks[hall.DbRows[i].Seats[j].Rank] = []Row{row}
					}
				}
				lastRank = hall.DbRows[i].Seats[j].Rank
				row = Row{
					IsRTL:  direction,
					Number: hall.DbRows[i].Number,
					Seats:  make([]Seat, 0),
				}
			}

			if hall.DbRows[i].Seats[j].State == 0 {
				row.Seats = append(row.Seats, hall.DbRows[i].Seats[j])
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
	db.Model(&seat).Updates(Seat{State: 1, UserID: userID})
}

func GetUser(userID uint) User {
	var user User
	user.ID = userID
	db.First(&user)
	var s []Seat
	db.Model(&Seat{}).Where("user_id = ?", userID).Find(&s)
	user.Seats = s

	return user
}

func CreateUser(name string) User {
	user := User{
		Name: name,
	}
	db.Create(&user)
	return user
}
