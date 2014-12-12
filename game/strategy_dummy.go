package game

type DummyStrategy struct{}

func (strategy *DummyStrategy) Reset() {
	// Nothing to reset.
}

// The dummy strategy actually works by keeping all the dice. Always.
func (strategy *DummyStrategy) ChooseDice(roll []int) []int {
	return roll
}
