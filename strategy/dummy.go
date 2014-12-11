package strategy

type DummyStrategy struct{}

// The dummy strategy actually works by keeping all the dice. Always.
func (strategy DummyStrategy) ChooseDice(roll []int) []int {
	return roll
}
