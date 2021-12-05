package tradestats

import "math"

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

func (stats *TradeStats) Score() float64 { // TODO: Take care of zero division
	return float64(stats.wins) / float64(stats.losses) * (100 / (1 + math.Pow(math.E, -1*float64(stats.matches))))
}
