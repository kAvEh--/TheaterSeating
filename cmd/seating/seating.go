package seating

import "github.com/kAvEh--/TheaterSeating/cmd/model"

func FindBestAlgorithm(h model.Hall, rank string, users []int) *model.Hall {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)

	h1, _ := deepCopy(h)
	h2, _ := deepCopy(h)
	h3, _ := deepCopy(h)

	go FillSeating(h1, rank, users, c1)
	go FillSeatsBestFit(h2, rank, users, findBestFit, c2)
	go FillSeatsBestFit(h3, rank, users, findFirstFit, c3)

	b1, b2, b3 := 0, 0, 0
	for i := 0; i < 3; i++ {
		select {
		case b1 = <-c1:
			if b1 == 0 {
				return &h1
			}
		case b2 = <-c2:
			if b2 == 0 {
				return &h2
			}
		case b3 = <-c3:
			if b3 == 0 {
				return &h3
			}
		}
	}

	if b2 <= b1 && b2 <= b3 {
		return &h2
	}
	if b3 <= b1 && b3 <= b2 {
		return &h3
	}

	return &h1
}

func FillSeating(h model.Hall, rank string, users []int, c chan int) {
	rIndex := 0
	sIndex := 0
	isLtr := true
	r := h.Ranks[rank]
	breakNum := 0
	for i := 0; i < len(users); i++ {
		for j := 0; j < users[i]; j++ {
			if r[rIndex].Seats[sIndex].State == 0 {
				r[rIndex].Seats[sIndex].State = 1
				r[rIndex].Seats[sIndex].User = i + 1
				if isLtr {
					sIndex++
					if sIndex >= len(r[rIndex].Seats) && rIndex < len(r)-1 {
						rIndex++
						sIndex = len(r[rIndex].Seats) - 1
						isLtr = !isLtr
						if j < users[i]-1 {
							breakNum++
						}
					}
				} else {
					sIndex--
					if sIndex < 0 && rIndex < len(r)-1 {
						rIndex++
						sIndex = 0
						isLtr = !isLtr
						if j < users[i]-1 {
							breakNum++
						}
					}
				}
			}
		}
	}
	c <- breakNum
}

func FillSeatsBestFit(h model.Hall, rank string, users []int, rowFinder func(model.Hall, string, int) int, c chan int) {
	rIndex := 0
	breakNum := 0
	for i := 0; i < len(users); i++ {
		rIndex = rowFinder(h, rank, users[i])
		if rIndex == -1 {
			breakNum++
			tmp := users[i]
			for tmp > 0 {
				j := 0
				for h.Ranks[rank][j].EmptySeatNum == 0 {
					j++
				}
				min := min(h.Ranks[rank][j].EmptySeatNum, tmp)
				tmp -= h.Ranks[rank][j].EmptySeatNum
				reserve(&h, rank, j, i+1, min)
			}
		} else {
			reserve(&h, rank, rIndex, i+1, users[i])
		}
	}

	c <- breakNum
}

func reserve(h *model.Hall, rank string, row int, group int, count int) {
	if h.Ranks[rank][row].IsRTL {
		index := len(h.Ranks[rank][row].Seats) - 1
		for i := 0; i < count; i++ {
			for h.Ranks[rank][row].Seats[index].State != 0 {
				index--
			}
			h.Ranks[rank][row].Seats[index].State = 1
			h.Ranks[rank][row].Seats[index].User = group
			h.Ranks[rank][row].EmptySeatNum--
		}
	} else {
		index := 0
		for i := 0; i < count; i++ {
			for h.Ranks[rank][row].Seats[index].State != 0 {
				index++
			}
			h.Ranks[rank][row].Seats[index].State = 1
			h.Ranks[rank][row].Seats[index].User = group
			h.Ranks[rank][row].EmptySeatNum--
		}
	}
}

func findBestFit(h model.Hall, rank string, num int) int {
	ret := -1
	for i := 0; i < len(h.Ranks[rank]); i++ {
		if h.Ranks[rank][i].EmptySeatNum >= num {
			if ret == -1 || h.Ranks[rank][i].EmptySeatNum < h.Ranks[rank][ret].EmptySeatNum {
				ret = i
			}
		}
	}

	return ret
}

func findFirstFit(h model.Hall, rank string, num int) int {
	for i := 0; i < len(h.Ranks[rank]); i++ {
		if h.Ranks[rank][i].EmptySeatNum >= num {
			return i
		}
	}

	return -1
}
