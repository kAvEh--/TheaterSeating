package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string
	Seats []Seat
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
	Seats  []Seat
}

type DBHall struct {
	gorm.Model
	Name string
	Rows []Row
}

type Hall struct {
	ID    uint
	Name  string
	Ranks map[string][]Row
}

type Row struct {
	ID           uint
	Number       int
	Seats        []Seat
	EmptySeatNum int
	IsRTL        bool
}
