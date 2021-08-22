package seating

func (h *Hall) FillSeating(rank string, users []int) {
	rIndex := 0
	sIndex := 0
	isLtr := true
	for i := 0; i < len(users); i++ {
		for j := 0; j < users[i]; j++ {
			if h.Ranks[rank].Rows[rIndex].Seats[sIndex].State == 0 {
				h.Ranks[rank].Rows[rIndex].Seats[sIndex].State = 1
				h.Ranks[rank].Rows[rIndex].Seats[sIndex].User = i + 1
				sIndex++
				if sIndex >= len(h.Ranks[rank].Rows[rIndex].Seats) {
					sIndex = 0
					rIndex++
					isLtr = !isLtr
				}
			}
		}
	}
}
