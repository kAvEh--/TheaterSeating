package seating

type Hall struct {
	ID    int
	Name  string
	Ranks map[string]Rank
}

type Rank struct {
	ID   int
	Name string
	Rows []Row
}

type Row struct {
	ID     int
	Number int
	Seats  []Seat
}

type Seat struct {
	ID    int
	State int
	User  int
}
