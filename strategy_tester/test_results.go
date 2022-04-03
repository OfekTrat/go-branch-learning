package strategytester

import st "branch_learning/strategy"

type TestResults struct {
	Wins           int
	Losses         int
	ConditionCount int
	StopLoss       float64
	TakeProfit     float64
	Score          float64
}

func CreateTestResultsFromStrategy(strategy *st.Strategy) *TestResults {
	return &TestResults{
		Wins:           0,
		Losses:         0,
		ConditionCount: strategy.Conditions().Length(),
		StopLoss:       strategy.StopLoss(),
		TakeProfit:     strategy.TakeProfit(),
		Score:          0,
	}
}

func (tr *TestResults) AddWins(wins int) {
	tr.Wins += wins
}

func (tr *TestResults) AddLosses(losses int) {
	tr.Losses += losses
}

func (tr *TestResults) CalcScore() {
	tr.Score = Score(tr)
}
