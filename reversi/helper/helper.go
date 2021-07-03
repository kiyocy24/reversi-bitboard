package helper

import (
	"errors"
	"github.com/kiyocy24/reversi-bitboard/reversi/board"
	"log"
	"strings"
)

func CoordinateToBit(row, col int) (bi uint64) {
	shift := (board.Length-row-1)*board.Length + (board.Length - col) - 1
	return 1 << shift
}

func BitToCoordinate(bi uint64) (int, int) {
	var index int
	for i := 0; i < board.AreaSize; i++ {
		if 1<<i == bi {
			index = i
		}
	}
	index = board.AreaSize - index - 1
	row := index / board.Length
	col := index % board.Length
	return row, col
}

func StringToCoordinate(s string) (row, col int, err error) {
	runes := []rune(strings.ToUpper(s))
	if len(runes) != 2 {
		return 0, 0, errors.New("length is not 2")
	}

	// row
	if runes[1] < '1' || '1'+board.Length-1 < runes[1] {
		log.Println("input", runes[1])
		return 0, 0, errors.New("invalid row arg")
	}
	for i := 0; i < board.Length; i++ {
		if runes[1] == rune('1'+i) {
			row = i
			break
		}
	}

	// col
	if runes[0] < 'A' || 'A'+board.Length-1 < runes[0] {
		return 0, 0, errors.New("invalid column arg")
	}
	for i := 0; i < board.Length; i++ {
		if runes[0] == rune('A'+i) {
			col = i
			break
		}
	}

	return row, col, err
}
