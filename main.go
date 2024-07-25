package main

import (
	"fmt"
)

func main() {
	board  := CreateBoard()
	fmt.Println(board)
	currentplayer := "X"
	for {
		PrintBoard(board)
		row , col := GetPoints(currentplayer)
		board[row][col] = currentplayer
		if CheckWin(board, currentplayer) {
			PrintBoard(board)
			fmt.Printf("Player %v Wins!!!", currentplayer)
			break
		}
		
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

func GetPoints(currentplayer string) (int, int){
	var col , row int
	for {
		fmt.Printf("Its Players %v turn\n", currentplayer)
		fmt.Println("Choose the column:")
		fmt.Scan(&col)
		fmt.Println("Choose the Row:")
		fmt.Scan(&row)
		if (row == 1 || row == 2 || row == 3) &&  (col == 1 || col == 2 || col == 3) {
			break
		}
		fmt.Println("Wrong choice of input Please choose again")
	}
	return row, col
}

func CheckWin(board [][]string, player string) bool {
	if board[1][1] == player && board[1][2] == player && board[1][3] == player {
		return true
	}
	if board[2][1] == player && board[2][2] == player && board[2][3] == player {
		return true
	}
	if board[3][1] == player && board[3][2] == player && board[3][3] == player {
		return true
	}
	if board[1][1] == player && board[2][2] == player && board[3][3] == player {
		return true
	}
	if board[3][1] == player && board[2][2] == player && board[1][3] == player {
		return true
	}
	return false
} 