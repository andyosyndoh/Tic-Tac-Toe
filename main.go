package main

import (
	"fmt"
	"os"
)

func main() {
	var mode string
	var attempts int
	fmt.Println("\033[2J\033[H")
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
		row, col := GetPoints(currentplayer, board)
		board[row][col] = currentplayer
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
			row, col = GetPoints(currentplayer, board)
		} else if currentplayer == "0" {
			fmt.Printf("Computer's Turn\n")
			row, col = Getcomputer(board)
		}
		board[row][col] = currentplayer
		if CheckWin(board, currentplayer) {
			// fmt.Println("\033[2J\033[H")
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
				fmt.Println("close win")
				return row, col
			} else {
				row, col = Place(temp)
				fmt.Println("place")
				return row, col
			}
		}
	}
}

func Mywin(temp [][]string) (int, int, error) {
	for i := range temp {
		for j := range temp[i] {
			if temp[i][j] == " " {
				temp[i][j] = "0"
				if CheckWin(temp, "0") {
					return j, i, nil
				}
				temp[i][j] = " "
			}
		}
	}
	return 0, 0, fmt.Errorf("not valid")
}

func OppWin(temp [][]string) (int, int, error) {
	for i := range temp {
		for j := range temp[i] {
			if temp[i][j] == " " {
				temp[i][j] = "X"
				if CheckWin(temp, "X") {
					return j, i, nil
				}
				temp[i][j] = " "
			}
		}
	}
	return 0, 0, fmt.Errorf("not valid")
}

func CloseWin(temp [][]string) (int, int, error) {
	for i := range temp {
		for j := range temp[i] {
			if temp[i][j] == " " {
				temp[i][j] = "0"
				countConnections := 0
				if i > 0 && temp[i-1][j] == "0" {
					countConnections++
				}
				if i < len(temp)-1 && temp[i+1][j] == "0" {
					countConnections++
				}
				if j > 0 && temp[i][j-1] == "0" {
					countConnections++
				}
				if j < len(temp[i])-1 && temp[i][j+1] == "0" {
					countConnections++
				}
				if countConnections > 0 {
					return j, i, nil
				}
				temp[i][j] = " "
			}
		}
	}
	return 0, 0, fmt.Errorf("not valid")
}

func Place(board [][]string) (int, int) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == " " {
				return i, j
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
		}
		fmt.Println()

	}
}

func GetPoints(currentplayer string, board [][]string) (int, int) {
	var col, row int
	for {
		fmt.Println("Choose the column:")
		fmt.Scan(&col)
		fmt.Println("Choose the Row:")
		fmt.Scan(&row)
		if (row == 1 || row == 2 || row == 3) && (col == 1 || col == 2 || col == 3) && board[row][col] == " " {
			break
		}
		fmt.Println("\033[2J\033[H")
		PrintBoard(board)
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
	} else {
		return "X"
	}
}
