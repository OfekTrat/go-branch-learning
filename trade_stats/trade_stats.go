package tradestats

type TradeStats struct {
	wins   int
	losses int
}

func (stats *TradeStats) Wins() int {
	return stats.wins
}
func (stats *TradeStats) Losses() int {
	return stats.losses
}

func (stats *TradeStats) AddWin() {
	stats.wins++
}
func (stats *TradeStats) AddLoss() {
	stats.losses++
}

func (stats *TradeStats) Score() float64 {
	var winLoseRatio float64

	if stats.losses == 0 {
		winLoseRatio = float64(stats.wins) / 0.1
	} else {
		winLoseRatio = float64(stats.wins) / float64(stats.losses)
	}
	return winLoseRatio * float64(stats.wins+stats.losses)
}
