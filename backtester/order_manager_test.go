package backtester

import (
	"branch_learning/exit"
	"testing"
)

func TestOrderManager_AddExit(t *testing.T) {
	om := &OrderManager{}
	exit1 := exit.CreateExit(100, 80)
	exit2 := exit.CreateExit(90, 70)

	om.AddExit(exit1)
	om.AddExit(exit2)
	if om.Exits()[0] != exit1 {
		t.Logf("Expected: %v\tGot: %v", exit1, om.Exits()[0])
		t.Error("AssertionError")
	}
	if om.Exits()[1] != exit2 {
		t.Logf("Expected: %v\tGot: %v", exit2, om.Exits()[1])
		t.Error("AssertionError")
	}

	if len(om.Exits()) != 2 {
		t.Logf("Expected Length: %v\tGot: %v", 2, len(om.Exits()))
		t.Error("AssertionError")
	}
}

func TestOrderManager_CheckingExists(t *testing.T) {
	om := &OrderManager{}
	exit1 := exit.CreateExit(100, 80)
	exit2 := exit.CreateExit(90, 70)
	om.AddExit(exit1)
	om.AddExit(exit2)

	wins, losses := om.CheckExits(85, 74)

	if wins != 0 {
		t.Logf("Expected: %v\tGot: %v", 0, wins)
		t.Error("AssertionError")
	}
	if losses != 1 {
		t.Logf("Expected: %v\tGot: %v", 1, losses)
		t.Error("AssertionError")
	}

	exits := om.Exits()

	if len(exits) != 1 {
		t.Log(exits)
		t.Logf("Expected Length: %v\tGot: %v", 1, len(exits))
		t.Error("AssertionError")
	}
	if exits[0] != exit2 {
		t.Logf("Expected: %v\tGot: %v", exit2, exits[0])
		t.Error("AssertionError")
	}
}
