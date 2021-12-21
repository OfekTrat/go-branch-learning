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
