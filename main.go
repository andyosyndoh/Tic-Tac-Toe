package main

import (
	"fmt"
)

func main() {
	board  := CreateBoard()
	fmt.Println(board)
	// currentplayer := "X"
	for {
		PrintBoard(board)
		break
	}
}

func CreateBoard() [][]string {
	board := make([][]string, 4)
	for i := range board {
		board[i] = make([]string, 4)
		for j := range board[i] {
			board[i][j] = " "
		}
	}

	board[0][0] = "🙂"
	board[0][1] = "1"
	board[0][2] = "2"
	board[0][3] = "3"
	board[1][0] = "1"
	board[2][0] = "2"
	board[3][0] = "3"

	return board
}

func PrintBoard(board [][]string){
	for _, line := range board {
		for _, str := range line {
			fmt.Print(str)
			fmt.Print(" ")
		}
		fmt.Println()
		
	}
}
