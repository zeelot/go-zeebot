package game

type Strategy interface {
	Reset()
	ChooseDice(roll []int) []int
}
