package condition

import (
	candle "branch_learning/candle"
	"testing"
)

func TestRedCondition_falseCase(t *testing.T) {
	c := candle.CreateCandle(1.1, 2.2, 3.3, 4.4)
	candleList := []candle.Candle{c}
	redCondition := RedCondition{CandleIndex: 0}

	answer := redCondition.DoesApply(&stream)

	if answer {
		t.Logf("Got: %v\tExpected %v", answer, false)
		t.Error("Assertion Error")
	}
}
