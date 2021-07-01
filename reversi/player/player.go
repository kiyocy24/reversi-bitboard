package player

type Player int

const (
	None Player = iota
	Black
	White
)

func (p Player) IsBlack() bool {
	return p == Black
}

func (p Player) IsWhite() bool {
	return p == White
}

func (p Player) String() string {
	if p.IsBlack() {
		return "Black"
	} else if p.IsWhite() {
		return "White"
	}

	return ""
}
