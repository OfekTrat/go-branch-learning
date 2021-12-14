package evolutioner

import "testing"

func TestChanceCreation(t *testing.T) {
	btScores := []float64{4, 10, 23, 9, 2, 6, 5}
	scs := createScores(btScores)
	chs := createChances(scs)

	if chs[0].strategyIndex != 4 {
		t.Error("AssertionError")
	}
	if chs[1].score != 6 {
		t.Logf("Expected 6 Got %f", chs[1].score)
		t.Error("AssertionError")
	}
	if chs[len(chs)-1].score != 59 && chs[len(chs)-1].strategyIndex != 2 {
		t.Error("AssertionError")
	}
}

func TestChance_GetByChance(t *testing.T) {
	btScores := []float64{4, 10, 23, 9, 2, 6, 5}
	chs := calcChances(btScores)

	i1 := chs.getIndexByChance(58)
	i2 := chs.getIndexByChance(0)
	i3 := chs.getIndexByChance(10)

	if i1 != 2 {
		t.Logf("Expected %v Got %v", 2, i1)
		t.Error("AssertionError")
	}
	if i2 != 4 {
		t.Logf("Expected %v Got %v", 4, i2)
		t.Error("AssertionError")
	}
	if i3 != 6 {
		t.Logf("Expected %v Got %v", 6, i3)
		t.Error("AssertionError")
	}
}

func TestScoresCreation(t *testing.T) {
	btScores := []float64{4, 10, 23, 9, 2, 6, 5}
	scs := createScores(btScores)

	for i, s := range btScores {
		if scs[i].score != s {
			t.Error("AssertionError")
		}
	}
}
