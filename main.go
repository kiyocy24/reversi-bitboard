package main

import (
	"fmt"
	"log"

	"github.com/kiyocy24/reversi-bitboard/reversi/board"
)

func main() {
	log.Println("start")

	b := board.NewBoard()
	var row, col int
	for {
		player := ""
		if b.Player().IsBlack() {
			player = "black"
		} else if b.Player().IsWhite() {
			player = "white"
		}
		fmt.Printf("[%s] (col, row) > ", player)
		_, err := fmt.Scan(&row, &col)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
