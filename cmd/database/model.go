package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Seats []Seat `gorm:"-"`
}

type Seat struct {
	gorm.Model `json:"-"`
	Rank       string
	State      int
	UserID     uint
	IsAisle    bool
	IsFrontRow bool
	IsHighSeat bool
	DBRowID    uint `json:"-"`
}

type DBRow struct {
	gorm.Model `json:"-"`
	Number     int
	Seats      []Seat `gorm:"foreignkey:DBRowID;" json:"seats"`
	DBHallID   uint
}

type DBHall struct {
	gorm.Model
	Name   string
	DbRows []DBRow `gorm:"foreignkey:DBHallID;" json:"rows"`
}

type Hall struct {
	ID    uint
	Name  string
	Ranks map[string][]Row
}

type Row struct {
	gorm.Model   `json:"-"`
	Number       int
	Seats        []Seat
	EmptySeatNum int
	IsRTL        bool
}
