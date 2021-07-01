package board

import (
	"errors"
	"github.com/kiyocy24/reversi-bitboard/reversi/bit"
	"github.com/kiyocy24/reversi-bitboard/reversi/direction"
	"github.com/kiyocy24/reversi-bitboard/reversi/player"
)

const (
	Length   = 8
	AreaSize = Length * Length
)

type Board struct {
	black  *bit.Bit      // 黒番
	white  *bit.Bit      // 白番
	player player.Player // 手番
	turn   int           // 手数
}

func NewBoard() *Board {
	return &Board{
		black:  bit.NewBit(),
		white:  bit.NewBit(),
		player: player.Black,
		turn:   0,
	}
}

func (b *Board) Reverse(p player.Player, input uint64) error {
	if b.player != p {
		return errors.New("not turn player")
	}

	// ひっくり返す
	b.reverse(input)

	return nil
}

func (b *Board) Clone() *Board {
	clone := *b
	return &clone
}

func (b *Board) LegalBoard() uint64 {
	playerBoard := *b.PlayerBoard()
	oppositeBoard := *b.OppositeBoard()
	horizontalBoard := oppositeBoard & 0x7e7e7e7e7e7e7e7e
	verticalBoard := oppositeBoard & 0x00FFFFFFFFFFFF00
	allSideBoard := oppositeBoard & horizontalBoard & verticalBoard
	blankBoard := ^(playerBoard | oppositeBoard)

	var tmp uint64
	var legalBoard uint64

	// 左上
	tmp = allSideBoard & (playerBoard << 9)
	for i := 0; i < Length-2; i++ {
		tmp |= allSideBoard & (tmp << 9)
	}
	legalBoard = blankBoard & (tmp << 9)

	// 上
	tmp = horizontalBoard & (playerBoard << 8)
	for i := 0; i < Length-2; i++ {
		tmp |= horizontalBoard & (tmp << 8)
	}
	legalBoard |= blankBoard & (tmp << 8)

	// 右上
	tmp = allSideBoard & (playerBoard << 7)
	for i := 0; i < Length-2; i++ {
		tmp |= allSideBoard & (tmp << 7)
	}
	legalBoard |= blankBoard & (tmp << 7)

	// 左
	tmp = verticalBoard & (playerBoard << 1)
	for i := 0; i < Length-2; i++ {
		tmp |= verticalBoard & (tmp << 1)
	}
	legalBoard |= blankBoard & (tmp << 1)

	// 右
	tmp = verticalBoard & (playerBoard << 1)
	for i := 0; i < Length-2; i++ {
		tmp |= verticalBoard & (tmp << 1)
	}
	legalBoard |= blankBoard & (tmp << 1)

	// 左下
	tmp = allSideBoard & (playerBoard >> 7)
	for i := 0; i < Length-2; i++ {
		tmp |= allSideBoard & (tmp >> 7)
	}
	legalBoard = blankBoard & (tmp >> 7)

	// 下
	tmp = verticalBoard & (playerBoard >> 8)
	for i := 0; i < Length-2; i++ {
		tmp |= verticalBoard & (tmp >> 8)
	}
	legalBoard |= blankBoard & (tmp >> 8)

	// 右下
	tmp = allSideBoard & (playerBoard >> 9)
	for i := 0; i < Length-2; i++ {
		tmp |= allSideBoard & (tmp >> 9)
	}
	legalBoard |= blankBoard & (tmp >> 9)

	return legalBoard
}

func (b *Board) Player() player.Player {
	return b.player
}

func (b *Board) PlayerBoard() *uint64 {
	if b.player.IsBlack() {
		return &b.black.Value
	} else if b.player.IsWhite() {
		return &b.white.Value
	}

	return nil
}

func (b *Board) OppositeBoard() *uint64 {
	if b.player.IsBlack() {
		return &b.white.Value
	} else if b.player.IsWhite() {
		return &b.black.Value
	}

	return nil
}

func (b *Board) reverse(put uint64) {
	playerBoard := b.PlayerBoard()
	oppositeBoard := b.OppositeBoard()

	// 反転ボード
	rev := uint64(0)
	for _, d := range direction.GetDirections() {
		r := uint64(0)
		mask := transfer(put, d)
		for mask != 0 && (mask&*oppositeBoard) != 0 {
			r |= mask
			mask = transfer(mask, d)
		}
		if mask&*playerBoard != 0 {
			rev |= r
		}
	}

	// 反転
	*playerBoard ^= put | rev
	*oppositeBoard ^= rev

	// 手番を進める
	b.next()
}

func transfer(put uint64, d direction.Direction) uint64 {
	switch d {
	case direction.Up:
		return (put << Length) & 0xffffffffffffff00
	case direction.UpperRight:
		return (put << (Length - 1)) & 0x7f7f7f7f7f7f7f00
	case direction.Right:
		return (put >> 1) & 0x7f7f7f7f7f7f7f7f
	case direction.LowerRight:
		return (put >> (Length + 1)) & 0x007f7f7f7f7f7f7f
	case direction.Low:
		return (put >> Length) & 0x00ffffffffffffff
	case direction.LowerLeft:
		return (put >> (Length - 1)) & 0x00fefefefefefefe
	case direction.Left:
		return (put << 1) & 0xfefefefefefefefe
	case direction.UpperLeft:
		return (put << (Length + 1)) & 0xfefefefefefefe00
	}

	return 0
}

func (b *Board) get(bi uint64) player.Player {
	if (b.black.Get() & bi) != 0 {
		return player.Black
	}
	if (b.white.Get() & bi) != 0 {
		return player.White
	}

	return player.None
}

func (b *Board) next() {
	if b.player.IsBlack() {
		b.player = player.White
	} else if b.player.IsWhite() {
		b.player = player.Black
	}

	b.turn++
}
