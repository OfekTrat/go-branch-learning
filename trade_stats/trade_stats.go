package tradestats

type TradeStats struct {
	wins    int
	losses  int
	matches int
}

func (stats *TradeStats) Wins() int {
	return stats.wins
}
func (stats *TradeStats) Losses() int {
	return stats.losses
}
func (stats *TradeStats) Matches() int {
	return stats.matches
}

func (stats *TradeStats) AddWin() {
	stats.wins++
	stats.matches++
}
func (stats *TradeStats) AddLoss() {
	stats.losses++
	stats.matches++
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
