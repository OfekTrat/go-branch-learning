package evolutioner

import "sort"

type indexScorePair struct {
	strategyIndex int
	score         float64
}
type scores []indexScorePair

func (s scores) Len() int           { return len(s) }
func (s scores) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s scores) Less(i, j int) bool { return s[i].score < s[j].score }

type chances []indexScorePair

func (c chances) getIndexByChance(chance float64) int {
	for _, isp := range c {
		if isp.score > chance {
			return isp.strategyIndex
		}
	}
	return c[len(c)-1].strategyIndex
}

func calcChances(btScores []float64) chances {
	s := createScores(btScores)
	chs := createChances(s)
	return chs
}

func createScores(btScores []float64) scores {
	s := make(scores, len(btScores))

	for i, score := range btScores {
		s[i] = indexScorePair{strategyIndex: i, score: score}
	}
	return s
}

func createChances(s scores) chances {
	sort.Sort(s)
	chs := make(chances, len(s))
	currentSum := float64(0)
	for i, score := range s {
		currentSum += score.score
		chs[i] = indexScorePair{score.strategyIndex, score.score}
	}
	return chs
}
