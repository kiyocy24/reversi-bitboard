package main

import (
	"fmt"
	"github.com/kiyocy24/reversi-bitboard/reversi"
	"log"
)

func main() {
	log.Println("start")

	r := reversi.NewReversi()
	var row, col int
	err := func() error {
		for {
			player := r.Player()
			fmt.Println(r.GetBoard())
			fmt.Printf("[%s] (col, row) > ", player)
			_, err := fmt.Scan(&row, &col)
			if err != nil {
				return err
			}
			if row < 0 || reversi.Length <= row {
				continue
			}
			if col < 0 || reversi.Length <= col {
				continue
			}

			err = r.Reverse(row, col)
			if err != nil {
				return err
			}
		}
	}()

	if err != nil {
		panic(err)
	}
}
