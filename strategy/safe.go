package strategy

import (
	"sort"
)

type SafeStrategy struct {
	HaveOne  bool
	HaveFour bool
}

func (strategy *SafeStrategy) Reset() {
	strategy.HaveOne = false
	strategy.HaveFour = false
}

func (strategy *SafeStrategy) ChooseDice(roll []int) []int {
	var chosen []int
	if strategy.HaveOne == true && strategy.HaveFour == true {
		sort.Sort(sort.Reverse(sort.IntSlice(roll)))
		for _, value := range roll {
			if value > 4 {
				// Keep any 5s and 6s
				chosen = append(chosen, value)
			}
			if value == 4 && len(chosen) == 0 {
				// Keeps the 4s if we haven't kept any yet.
				chosen = append(chosen, 4)
			}
		}
		if len(chosen) == 0 {
			// Simply keep the highest one.
			chosen = append(chosen, roll[0])
		}
	} else {
		sort.Sort(sort.IntSlice(roll))
		for _, value := range roll {
			if value == 1 && strategy.HaveOne != true {
				chosen = append(chosen, 1)
				strategy.HaveOne = true
			}
			if value == 4 && strategy.HaveFour != true {
				chosen = append(chosen, 4)
				strategy.HaveFour = true
			}
		}
		if len(chosen) == 0 {
			// Simply keep the highest one.
			chosen = append(chosen, roll[len(roll)-1])
		}
	}

	return chosen
}
