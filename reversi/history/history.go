package history

import (
	"github.com/google/uuid"
	"github.com/kiyocy24/reversi-bitboard/reversi/board"
)

type History struct {
	Id       string
	Board    *board.Board
	nexts    []*History
	previous *History
}

func NewHistory(b *board.Board) *History {
	return &History{
		Id:       uuid.NewString(),
		Board:    b.Clone(),
		previous: nil,
	}
}

func (h *History) AddHistory(b *board.Board) *History {
	next := &History{
		Id:       uuid.NewString(),
		Board:    b.Clone(),
		previous: h,
	}
	h.nexts = append(h.nexts, next)
	return next
}
