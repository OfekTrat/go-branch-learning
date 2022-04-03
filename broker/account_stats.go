package broker

type AccountStats struct {
	wins   int
	losses int
}

func CreateEmptyAccountStats() *AccountStats {
	return &AccountStats{0, 0}
}

func (as *AccountStats) AddWin() {
	as.wins++
}

func (as *AccountStats) AddLoss() {
	as.losses++
}

func (as *AccountStats) Wins() int {
	return as.wins
}

func (as *AccountStats) Losses() int {
	return as.losses
}

func (as *AccountStats) AddAccountStats(other *AccountStats) {
	as.wins += other.wins
	as.losses += other.losses
}
