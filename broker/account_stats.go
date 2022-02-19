package broker

type accountStats struct {
	wins   int
	losses int
}

func AccountStats() *accountStats {
	return &accountStats{0, 0}
}

func (as *accountStats) AddWin() {
	as.wins++
}

func (as *accountStats) AddLoss() {
	as.losses++
}

func (as *accountStats) Wins() int {
	return as.wins
}

func (as *accountStats) Losses() int {
	return as.losses
}
