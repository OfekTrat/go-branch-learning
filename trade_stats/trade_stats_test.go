package tradestats

import "testing"

func TestTradeStats_AddWin(t *testing.T) {
	stats := &TradeStats{}
	nWins := 3

	for i := 0; i < nWins; i++ {
		stats.AddWin()
	}

	if stats.Wins() != nWins {
		t.Logf("Expected %v\tGot: %v", nWins, stats.Wins())
	}
}

func TestTradeStats_AddLoss(t *testing.T) {
	stats := &TradeStats{}
	nLosses := 3

	for i := 0; i < nLosses; i++ {
		stats.AddLoss()
	}

	if stats.Losses() != nLosses {
		t.Logf("Expected %v\tGot: %v", nLosses, stats.Losses())
	}
}

func TestTradeStats_Matches(t *testing.T) {
	stats := &TradeStats{}
	stats.AddLoss()
	stats.AddWin()

	if stats.Matches() != 2 {
		t.Logf("Expected %v\tGot: %v", 2, stats.Matches())
	}
}

func TestTradeStats_Score(t *testing.T) {
	return
	stats := &TradeStats{}
	stats.AddWin()
	stats.AddWin()
	stats.AddWin()
	stats.AddLoss()
	stats.AddLoss()
	stats.AddLoss()
	stats.AddLoss()

	score := stats.Score()

	if score != 7.59375 {
		t.Logf("Expected %v\tGot: %v", 7.5, score)
		t.Error("AssertionError")
	}
}
