package strategy

import (
	"fmt"
	"testing"
)

type MockRoll struct {
	Input  []int
	Output []int
}

func TestBehavior(t *testing.T) {
	games := [][]MockRoll{
		[]MockRoll{
			MockRoll{
				Input:  []int{1, 4, 6, 6, 6},
				Output: []int{1, 4},
			},
			MockRoll{
				Input:  []int{6, 6, 6},
				Output: []int{6, 6, 6},
			},
		},
		[]MockRoll{
			MockRoll{
				Input:  []int{1, 4, 6, 6, 6},
				Output: []int{1, 4},
			},
			MockRoll{
				Input:  []int{1, 4, 6, 3, 6},
				Output: []int{6, 6},
			},
			MockRoll{
				Input:  []int{1, 4, 3, 2, 1},
				Output: []int{4},
			},
		},
		[]MockRoll{
			MockRoll{
				Input:  []int{5, 5, 5, 5, 5, 5},
				Output: []int{5},
			},
			MockRoll{
				Input:  []int{6, 6, 6, 6, 6},
				Output: []int{6},
			},
			MockRoll{
				Input:  []int{3, 2, 3, 2},
				Output: []int{3},
			},
			MockRoll{
				Input:  []int{1, 6, 6},
				Output: []int{1},
			},
			MockRoll{
				Input:  []int{1, 4},
				Output: []int{4},
			},
		},
		[]MockRoll{
			MockRoll{
				Input:  []int{5, 6, 4, 1, 6, 6},
				Output: []int{1, 4},
			},
			MockRoll{
				Input:  []int{3, 2, 5, 4},
				Output: []int{5},
			},
			MockRoll{
				Input:  []int{6, 5, 6},
				Output: []int{6, 6, 5},
			},
		},
	}

	for gameId, rolls := range games {
		safeStrategy := SafeStrategy{}
		for rollId, roll := range rolls {
			result := safeStrategy.ChooseDice(roll.Input)

			fmt.Printf("rolled %v and chose %v\n", roll.Input, result)

			if !compare(result, roll.Output) {
				t.Fatalf("Roll did not result in the right choices(gameId %d rollId %d): %v is not %v", gameId, rollId, result, roll.Output)
			}
		}
	}
}

func compare(X, Y []int) bool {
	m := make(map[int]int)

	for _, y := range Y {
		m[y]++
	}

	for _, x := range X {
		m[x]--
	}

	for _, z := range m {
		if z != 0 {
			return false
		}
	}

	return true
}
