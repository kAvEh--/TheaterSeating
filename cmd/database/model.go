package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Seats []Seat `gorm:"foreignKey:ID"`
}

type Seat struct {
	gorm.Model
	Rank       string
	State      int
	User       int
	IsAisle    bool
	IsFrontRow bool
	IsHighSeat bool
}

type DBRow struct {
	gorm.Model
	Number int
	Seats  []Seat `gorm:"foreignKey:ID"`
}

type DBHall struct {
	gorm.Model
	Name string
	Rows []DBRow `json:"rows" gorm:"foreignKey:ID"`
}

type Hall struct {
	ID    uint
	Name  string
	Ranks map[string][]Row
}

type Row struct {
	gorm.Model
	Number       int
	Seats        []Seat `gorm:"foreignKey:ID"`
	EmptySeatNum int
	IsRTL        bool
}
