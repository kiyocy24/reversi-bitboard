package direction

type Direction int

const (
	Up Direction = iota + 1
	UpperRight
	Right
	LowerRight
	Low
	LowerLeft
	Left
	UpperLeft
)

func GetDirections() []Direction {
	return []Direction{
		Up,
		UpperRight,
		Right,
		LowerRight,
		Low,
		LowerLeft,
		Left,
		UpperLeft,
	}
}
