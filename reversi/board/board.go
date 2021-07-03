package board

import (
	"errors"
	"github.com/kiyocy24/reversi-bitboard/reversi/direction"
	"github.com/kiyocy24/reversi-bitboard/reversi/player"
)

const (
	Length   = 8
	AreaSize = Length * Length
)

type Board struct {
	black    uint64        // 黒番
	white    uint64        // 白番
	player   player.Player // 手番
	opposite player.Player // 待ち番
	turn     int           // 手数
}

func NewBoard() *Board {
	return &Board{
		black:    E4 | D5,
		white:    D4 | E5,
		player:   player.Black,
		opposite: player.White,
		turn:     1,
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
	allSideBoard := oppositeBoard & 0x007e7e7e7e7e7e00
	blankBoard := ^(playerBoard | oppositeBoard)

	var tmp uint64
	var legalBoard uint64

	// 上
	tmp = verticalBoard & (playerBoard << 8)
	tmp |= verticalBoard & (tmp << 8)
	tmp |= verticalBoard & (tmp << 8)
	tmp |= verticalBoard & (tmp << 8)
	tmp |= verticalBoard & (tmp << 8)
	tmp |= verticalBoard & (tmp << 8)
	legalBoard = blankBoard & (tmp << 8)

	// 下
	tmp = verticalBoard & (playerBoard >> 8)
	tmp |= verticalBoard & (tmp >> 8)
	tmp |= verticalBoard & (tmp >> 8)
	tmp |= verticalBoard & (tmp >> 8)
	tmp |= verticalBoard & (tmp >> 8)
	tmp |= verticalBoard & (tmp >> 8)
	legalBoard |= blankBoard & (tmp >> 8)

	// 左
	tmp = horizontalBoard & (playerBoard << 1)
	tmp |= verticalBoard & (tmp << 1)
	tmp |= verticalBoard & (tmp << 1)
	tmp |= verticalBoard & (tmp << 1)
	tmp |= verticalBoard & (tmp << 1)
	tmp |= verticalBoard & (tmp << 1)
	legalBoard |= blankBoard & (tmp << 1)

	// 右
	tmp = horizontalBoard & (playerBoard >> 1)
	tmp |= verticalBoard & (tmp >> 1)
	tmp |= verticalBoard & (tmp >> 1)
	tmp |= verticalBoard & (tmp >> 1)
	tmp |= verticalBoard & (tmp >> 1)
	tmp |= verticalBoard & (tmp >> 1)
	legalBoard |= blankBoard & (tmp >> 1)

	// 左上
	tmp = allSideBoard & (playerBoard << 9)
	tmp |= allSideBoard & (tmp << 9)
	tmp |= allSideBoard & (tmp << 9)
	tmp |= allSideBoard & (tmp << 9)
	tmp |= allSideBoard & (tmp << 9)
	tmp |= allSideBoard & (tmp << 9)
	legalBoard |= blankBoard & (tmp << 9)

	// 右上
	tmp = allSideBoard & (playerBoard << 7)
	tmp |= allSideBoard & (tmp << 7)
	tmp |= allSideBoard & (tmp << 7)
	tmp |= allSideBoard & (tmp << 7)
	tmp |= allSideBoard & (tmp << 7)
	tmp |= allSideBoard & (tmp << 7)
	legalBoard |= blankBoard & (tmp << 7)

	// 左下
	tmp = allSideBoard & (playerBoard >> 7)
	tmp |= allSideBoard & (tmp >> 7)
	tmp |= allSideBoard & (tmp >> 7)
	tmp |= allSideBoard & (tmp >> 7)
	tmp |= allSideBoard & (tmp >> 7)
	tmp |= allSideBoard & (tmp >> 7)
	legalBoard |= blankBoard & (tmp >> 7)

	// 右下
	tmp = allSideBoard & (playerBoard >> 9)
	tmp |= allSideBoard & (tmp >> 9)
	tmp |= allSideBoard & (tmp >> 9)
	tmp |= allSideBoard & (tmp >> 9)
	tmp |= allSideBoard & (tmp >> 9)
	tmp |= allSideBoard & (tmp >> 9)
	legalBoard |= blankBoard & (tmp >> 9)

	return legalBoard
}

func (b *Board) Player() player.Player {
	return b.player
}

func (b *Board) Opposite() player.Player {
	return b.opposite
}

func (b *Board) PlayerBoard() *uint64 {
	if b.player.IsBlack() {
		return &b.black
	} else if b.player.IsWhite() {
		return &b.white
	}

	return nil
}

func (b *Board) OppositeBoard() *uint64 {
	if b.player.IsBlack() {
		return &b.white
	} else if b.player.IsWhite() {
		return &b.black
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

func (b *Board) next() {
	if b.player.IsBlack() {
		b.player = player.White
		b.opposite = player.Black
	} else if b.player.IsWhite() {
		b.player = player.Black
		b.opposite = player.White
	}

	b.turn++
}
