package main

import (
	"fmt"
	"github.com/kiyocy24/reversi-bitboard/reversi"
	"github.com/kiyocy24/reversi-bitboard/reversi/helper"
	"log"
)

func main() {
	log.Println("start")

	r := reversi.NewReversi()
	var input string
	err := func() error {
		for {
			player := r.Player()
			fmt.Println(r.GetBoard())
			fmt.Printf("[%s] > ", player)
			_, err := fmt.Scan(&input)
			if err != nil {
				return err
			}

			row, col, err := helper.StringToCoordinate(input)
			if err != nil {
				fmt.Println(err)
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
