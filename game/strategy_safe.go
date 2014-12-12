package game

import (
	"sort"
)

type SafeStrategy struct {
	HaveOne  bool
	HaveFour bool
}

func NewSafeStrategy() *SafeStrategy {
	return &SafeStrategy{}
}

func (strategy *SafeStrategy) Reset() {
	strategy.HaveOne = false
	strategy.HaveFour = false
}

func (strategy *SafeStrategy) ChooseDice(roll []int) []int {
	chosen := map[int]int{}
	rollCopy := make([]int, len(roll), len(roll))
	copy(rollCopy, roll)
	sort.Sort(sort.Reverse(sort.IntSlice(rollCopy)))

	for k, v := range strategy.chooseQualifying(roll) {
		chosen[k] = v
	}

	for k, v := range strategy.chooseHighs(roll) {
		chosen[k] = v
	}

	if len(chosen) == 0 {
		// Simply keep the highest one.
		chosen[0] = rollCopy[0]
	}

	kept := []int{}
	for _, v := range chosen {
		kept = append(kept, v)
	}

	return kept
}

func (strategy *SafeStrategy) chooseQualifying(roll []int) map[int]int {
	chosen := make(map[int]int)
	if strategy.HaveOne && strategy.HaveFour {
		return chosen
	}

	for k, v := range roll {
		if v == 1 && !strategy.HaveOne {
			chosen[k] = 1
			strategy.HaveOne = true
		}
		if v == 4 && !strategy.HaveFour {
			chosen[k] = 4
			strategy.HaveFour = true
		}
	}

	return chosen
}

func (strategy *SafeStrategy) chooseHighs(roll []int) map[int]int {
	chosen := make(map[int]int)
	if !strategy.HaveOne || !strategy.HaveFour {
		return chosen
	}

	for k, v := range roll {
		if v > 4 {
			// Keep any 5s and 6s
			chosen[k] = v
		}
	}

	return chosen
}
