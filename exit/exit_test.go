package exit

import "testing"

func TestExit_StopLoss(t *testing.T) {
	exit := Exit{takeProfitPrice: 100, stopLossPrice: 80}
	if ans := exit.IsStop(70); !ans {
		t.Logf("Expected: %v\tGot: %v", true, ans)
		t.Error("AssertionError")
	}
}

func TestExit_TakeProfit(t *testing.T) {
	exit := Exit{takeProfitPrice: 100, stopLossPrice: 80}
	if ans := exit.IsTake(110); !ans {
		t.Logf("Expected: %v\tGot: %v", true, ans)
		t.Error("AssertionError")
	}
}

func TestExit_DoNothing(t *testing.T) {
	exit := Exit{takeProfitPrice: 100, stopLossPrice: 80}
	if ans := exit.IsStop(90) && exit.IsTake(90); ans {
		t.Logf("Expected: %v\tGot: %v", false, ans)
		t.Error("AssertionError")
	}
}
