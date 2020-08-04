package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/myoan/reversi"
)

func readPosition() *reversi.Position {
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Printf("X: ")
	stdin.Scan()
	xStr := stdin.Text()
	fmt.Printf("Y: ")
	stdin.Scan()
	yStr := stdin.Text()
	x, _ := strconv.Atoi(xStr)
	y, _ := strconv.Atoi(yStr)
	return &reversi.Position{X: x, Y: y}
}

func main() {
	game := reversi.NewGame()
	for {
		switch game.GameState {
		case reversi.Prepare:
			game.GameState = reversi.BlackTurn
		case reversi.BlackTurn:
			fmt.Println("Black turn")
			pos := readPosition()
			game.SetStone(1, pos)
		case reversi.WhiteTurn:
			fmt.Println("White turn")
			pos := readPosition()
			game.SetStone(2, pos)
		case reversi.Finish:
			fmt.Println("Finish")
			fmt.Printf("%d win!\n", game.Winner())
			return
		}
		game.Show()
	}
}
