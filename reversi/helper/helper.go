package helper

import "github.com/kiyocy24/reversi-bitboard/reversi/board"

func CoordinateToBit(row, col int) uint64 {
	var bi uint64
	bi = bi >> col
	bi = bi >> ((row - 1) * board.Length)

	return bi
}

func BitToCoordinate(bi uint64) (int, int) {
	var index int
	for i := 0; i < board.AreaSize; i++ {
		if 1<<i == bi {
			index = i
		}
	}

	row := index / board.Length
	col := index % board.Length
	return row, col
}
