package piscine

import "github.com/01-edu/z01"

const N = 8

// isSafe checks if a queen can be placed at board[row][col]
// This function is called when "col" queens are already placed in columns from 0 to col -1
func isSafe(board *[N]int, row, col int) bool {
	for i := 0; i < col; i++ {
		// Check this row on left side and check diagonals
		if board[i] == row || board[i] == row+(i-col) || board[i] == row-(i-col) {
			return false
		}
	}
	return true
}

// solve recursively places queens on the board
func solve(board *[N]int, col int) {
	if col == N {
		printBoard(board)
		return
	}

	for i := 0; i < N; i++ {
		if isSafe(board, i, col) {
			board[col] = i // Place the queen at the current position
			solve(board, col+1)
		}
	}
}

// printBoard prints the solution
// printBoard prints the solution
// printBoard prints the solution
func printBoard(board *[N]int) {
	for i := 0; i < N; i++ {
		z01.PrintRune(rune(board[i] + '1'))
	}
	z01.PrintRune('\n')
}

// EightQueens finds and prints all solutions to the eight queens puzzle
func EightQueens() {
	var board [N]int
	solve(&board, 0)
}
