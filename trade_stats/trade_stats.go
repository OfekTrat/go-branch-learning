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

func (stats *TradeStats) Score() float32 {
	return (float32(stats.wins) / float32(stats.losses)) * float32(stats.matches)
}
