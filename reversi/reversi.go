package reversi

import (
	"fmt"
	"github.com/kiyocy24/reversi-bitboard/reversi/board"
	"github.com/kiyocy24/reversi-bitboard/reversi/helper"
	"github.com/kiyocy24/reversi-bitboard/reversi/player"
)

const (
	Length   = board.Length
	AreaSize = board.AreaSize
	Black    = "X"
	White    = "O"
	Legal    = "+"
	NONE     = " "
)

type Reversi struct {
	b *board.Board
}

func NewReversi() *Reversi {
	return &Reversi{
		b: board.NewBoard(),
	}
}

func (r *Reversi) Player() string {
	if r.b.Player().IsBlack() {
		return Black
	} else if r.b.Player().IsWhite() {
		return White
	}

	return ""
}

func (r *Reversi) GetBoard() (s string) {
	bb := r.getBoard()
	legalBoard := r.getLegalBoard()

	s = " |"
	for i := 0; i < Length; i++ {
		s += fmt.Sprintf("%c|", 'A'+i)
	}
	s += "\n"

	for row := 0; row < Length; row++ {
		s += fmt.Sprintf("%d", row+1)
		for col := 0; col < Length; col++ {
			s += "|"
			b := bb[row][col]
			isLegal := legalBoard[row][col]
			if b.IsBlack() {
				s += Black
			} else if b.IsWhite() {
				s += White
			} else {
				if isLegal {
					s += Legal
				} else {
					s += NONE
				}
			}
		}
		s += "|\n"
	}

	return s
}

func (r *Reversi) Reverse(row, col int) error {
	err := r.b.Reverse(r.b.Player(), helper.CoordinateToBit(row, col))
	if err != nil {
		return err
	}

	return nil
}

func (r *Reversi) getBoard() [Length][Length]player.Player {
	bb := [Length][Length]player.Player{}
	playerBoard := *r.b.PlayerBoard()
	oppositeBoard := *r.b.OppositeBoard()

	bi := uint64(0)
	for i := 0; i < AreaSize; i++ {
		bi = 1 << i
		row, col := helper.BitToCoordinate(bi)
		if playerBoard&bi != 0 {
			bb[row][col] = r.b.Player()
		} else if oppositeBoard&bi != 0 {
			bb[row][col] = r.b.Opposite()
		} else {
			bb[row][col] = player.None
		}
	}

	return bb
}

func (r *Reversi) getLegalBoard() [Length][Length]bool {
	bb := [Length][Length]bool{}
	legalBoard := r.b.LegalBoard()

	bi := uint64(0)
	for i := 0; i < AreaSize; i++ {
		bi = bi << i
		row, col := helper.BitToCoordinate(bi)
		bb[row][col] = legalBoard&bi != 0
	}

	return bb
}
