package main

import (
	"fmt"
	"os"
)

func main() {

	var mode string
	var attempts int

	fmt.Println("Welcome To Tic-tac-toe,Please choose your preffered mode:\n1.Two Players\n2.One Player")
	for {
		fmt.Scan(&mode)
		if mode == "1" {
			TwoPlayers()
		} else if mode == "2" {
			OnePlayer()
		}
		if attempts == 4 {
			fmt.Println("Sorry: Maximum number of attempts")
			os.Exit(0)
		}
		fmt.Println("Wrong Choice.Please Choose Again")
		attempts++
	}

}

func TwoPlayers() {
	board := CreateBoard()
	currentplayer := "X"
	for {
		fmt.Println("\033[2J\033[H")
		PrintBoard(board)
		fmt.Printf("Its Players %v turn\n", currentplayer)
		row, col := GetPoints(currentplayer)
		board[row][col] = currentplayer + " "
		if CheckWin(board, currentplayer) {
			fmt.Println("\033[2J\033[H")
			PrintBoard(board)
			fmt.Printf("Player %v Wins!!!\n", currentplayer)
			os.Exit(0)
		}
		if canContinue(board) {
			currentplayer = ChangePlayer(currentplayer)
		} else {
			fmt.Println("Game Over")
			os.Exit(0)
		}
	}
}

func OnePlayer() {
	board := CreateBoard()
	currentplayer := "X"
	row, col := 0, 0
	for {
		// fmt.Println("\033[2J\033[H")
		PrintBoard(board)
		if currentplayer == "X" {
			fmt.Printf("Your Turn\n")
			row, col = GetPoints(currentplayer)
		} else if currentplayer == "0" {
			fmt.Printf("Computer's Turn\n")
			row, col = Getcomputer(board)
		}
		board[row][col] = currentplayer + " "
		if CheckWin(board, currentplayer) {
			// fmt.Println("\033[2J\033[H")
			PrintBoard(board)
			fmt.Printf("Player %v Wins!!!\n", currentplayer)
			os.Exit(0)
		}
		// if canContinue(board) {
		currentplayer = ChangePlayer(currentplayer)
		// } else {
		// 	fmt.Println("Game Over")
		// 	os.Exit(0)
		// }

	}
}

func canContinue(board [][]string) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == " " {
				return true
			}
		}
	}
	return false
}

func Getcomputer(board [][]string) (int, int) {

	temp := board
	row, col, err := Mywin(temp)
	if err == nil {
		fmt.Println("mine")
		return row, col
	} else {
		row, col, err = OppWin(temp)
		if err == nil {
			fmt.Println("opp")
			return row, col
		} else {
			row, col, err = CloseWin(temp)
			if err == nil {
				fmt.Println("close")
				return row, col
			} else {
				row, col = Place(temp)
				return row, col
			}
		}
	}
}

func Mywin(board [][]string) (int, int, error) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == " " {
				board[i][j] = "0"
				if CheckWin(board, "0") {
					return j, i, nil
				}
				board[i][j] = " "
			}
		}
	}
	return 0, 0, fmt.Errorf("not valid")
}

func OppWin(board [][]string) (int, int, error) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == " " {
				board[i][j] = "X"
				if CheckWin(board, "X") {
					return j, i, nil
				}
				board[i][j] = " "
			}
		}
	}
	return 0, 0, fmt.Errorf("not valid")
}

func CloseWin(board [][]string) (int, int, error) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == " " {
				
			}
		}
	}
	return 0, 0, fmt.Errorf("not valid")
}

func Place(board [][]string) (int, int) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == " " {
				return j, i
			}
		}
	}
	return 0, 0
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

func PrintBoard(board [][]string) {
	for _, line := range board {
		for _, str := range line {
			fmt.Print(str)
			fmt.Print(" ")
		}
		fmt.Println()

	}
}

func GetPoints(currentplayer string) (int, int) {
	var col, row int
	for {
		fmt.Println("Choose the column:")
		fmt.Scan(&col)
		fmt.Println("Choose the Row:")
		fmt.Scan(&row)
		if (row == 1 || row == 2 || row == 3) && (col == 1 || col == 2 || col == 3) {
			break
		}
		fmt.Println("Wrong choice of input Please choose again")
		row, col = 0, 0
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
	if board[1][1] == player && board[2][1] == player && board[3][1] == player {
		return true
	}
	if board[1][2] == player && board[2][2] == player && board[3][2] == player {
		return true
	}
	if board[1][3] == player && board[2][3] == player && board[3][3] == player {
		return true
	}
	return false
}

func ChangePlayer(c string) string {
	if c == "X" {
		return "0"
	}
	return "X"
}
