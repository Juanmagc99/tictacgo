package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

const min_symbol string = "X"
const max_symbol string = "O"

type Move struct {
	row    int
	column int
	score  int
}

func main() {

	board := [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}

	fmt.Println("Welcome to Tic-Tac-Toe!")
	PrintPet()
	PrintBoard(board)
	for {
		row, col := ChoosePosition(board)
		board[row][col] = min_symbol
		ClearConsole()
		PrintBoard(board)
		if CheckEnd(board) {
			break
		}
		move := MakeMove(board)
		board[move.row][move.column] = max_symbol
		ClearConsole()
		PrintBoard(board)
		if CheckEnd(board) {
			break
		}
	}

}

func CheckEnd(board [3][3]string) bool {
	isEnded := false
	score := EvaluateBoard(board)
	if score == -10 {
		fmt.Println("Congrats you have won!")
		isEnded = true
	} else if score == 10 {
		fmt.Println("Yo have lost!")
		isEnded = true
	} else if score == 0 && !MovesLeft(board) {
		fmt.Println("Its a draw!")
		isEnded = true
	}

	return isEnded
}

func MakeMove(board [3][3]string) Move {

	best_move := Move{
		row:    -1,
		column: -1,
		score:  -1000,
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				board[i][j] = max_symbol
				move := Move{row: i, column: j, score: MinMax(board, min_symbol)}
				board[i][j] = " "
				if best_move.score < move.score {
					best_move = move
				}
			}
		}
	}

	return best_move
}

func ChoosePosition(board [3][3]string) (int, int) {

	var row, col int

	fmt.Print("Choose position: ")
	_, err := fmt.Scan(&row, &col)
	if err != nil {
		fmt.Println("Error try again", err)
		ChoosePosition(board)
	} else if row > 2 || col > 2 || row < 0 || col < 0 {
		fmt.Println("Positions can be from (0,0) to (2,2)")
		ChoosePosition(board)
	} else if board[row][col] != " " {
		fmt.Println("That position is not empty!")
		ChoosePosition(board)
	}

	return row, col
}

func MinMax(board [3][3]string, player string) int {

	var score int = EvaluateBoard(board)

	if score == 10 || score == -10 {
		return score
	} else if !MovesLeft(board) {
		return 0
	}

	var best int

	if player == min_symbol {
		best = 100
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == " " {
					board[i][j] = min_symbol
					best = min(best, MinMax(board, max_symbol))
					board[i][j] = " "
				}
			}
		}
	} else {
		best = -100
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == " " {
					board[i][j] = max_symbol
					best = max(best, MinMax(board, min_symbol))
					board[i][j] = " "
				}
			}
		}
	}

	return best
}

func EvaluateBoard(board [3][3]string) int {

	/*
		Return -10 if Min win +10 if Max win and 0 if
		its a draw or there is no winner yet
	*/

	//Evaluate Rows and Colums
	//Evaluate Colums
	for index := range []int{0, 1, 2} {
		if board[index][0] == board[index][1] && board[index][0] == board[index][2] &&
			board[index][0] != " " {
			if board[index][0] == min_symbol {
				return -10
			} else {
				return 10
			}
		}

		if board[0][index] == board[1][index] && board[0][index] == board[2][index] &&
			board[0][index] != " " {
			if board[0][index] == min_symbol {
				return -10
			} else {
				return 10
			}
		}
	}

	//Evaluate Diagonals
	/*
		Hay que cambiar ya que 1 1 no tiene por que esta puesto todavia y pueden ser las dos esquinas
	*/
	if board[0][0] == board[1][1] && board[0][0] == board[2][2] && board[0][0] != " " {
		if board[1][1] == min_symbol {
			return -10
		} else {
			return 10
		}
	}

	if board[0][2] == board[1][1] && board[0][2] == board[2][0] && board[0][2] != " " {
		if board[1][1] == min_symbol {
			return -10
		} else {
			return 10
		}
	}

	return 0
}

func MovesLeft(board [3][3]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return true
			}
		}
	}
	return false
}

func PrintBoard(board [3][3]string) {
	for _, row := range board {
		fmt.Println("  =   =   =")
		for _, val := range row {
			fmt.Printf("| %s ", val)
		}
		fmt.Println("|")
	}
	fmt.Println("  =   =   =")
}

func PrintPet() {
	fmt.Println("   /\\_/\\  ")
	fmt.Println("  ( o o )  ")
	fmt.Println("   ( ^ )   ")
	fmt.Println("    '-'     ")
}

func ClearConsole() {
	so := runtime.GOOS
	if so == "linux" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
