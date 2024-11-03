package main

import (
	"fmt"
	"os"
)

const SIZE = 9

// Sudoku board
var board [SIZE][SIZE]int

// Solution board to store the first valid solution
var solution [SIZE][SIZE]int

func main() {
	// Validate input arguments
	if len(os.Args) != SIZE+1 {
		fmt.Println("Error")
		return
	}

	// Parse input into the board
	for i := 0; i < SIZE; i++ {
		line := os.Args[i+1]
		if len(line) != SIZE {
			fmt.Println("Error")
			return
		}
		for j, ch := range line {
			if ch == '.' {
				board[i][j] = 0
			} else if ch >= '1' && ch <= '9' {
				board[i][j] = int(ch - '0')
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	// Check if the initial board is valid
	if !isValid() {
		fmt.Println("Error")
		return
	}

	// Solve the Sudoku and check for uniqueness
	solutions := 0
	solve(&solutions, 2)

	if solutions == 1 {
		printBoard(solution)
	} else {
		fmt.Println("Error")
	}
}

// Check if the current board is valid
func isValid() bool {
	// Check rows and columns
	for i := 0; i < SIZE; i++ {
		row := make([]bool, SIZE+1)
		col := make([]bool, SIZE+1)
		for j := 0; j < SIZE; j++ {
			// Check row
			if board[i][j] != 0 {
				if row[board[i][j]] {
					return false
				}
				row[board[i][j]] = true
			}
			// Check column
			if board[j][i] != 0 {
				if col[board[j][i]] {
					return false
				}
				col[board[j][i]] = true
			}
		}
	}

	// Check 3x3 subgrids
	for i := 0; i < SIZE; i += 3 {
		for j := 0; j < SIZE; j += 3 {
			if !isValidSubgrid(i, j) {
				return false
			}
		}
	}

	return true
}

// Check if a 3x3 subgrid is valid
func isValidSubgrid(startRow, startCol int) bool {
	marks := make([]bool, SIZE+1)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			val := board[startRow+i][startCol+j]
			if val != 0 {
				if marks[val] {
					return false
				}
				marks[val] = true
			}
		}
	}
	return true
}

// Recursive solver to find all solutions up to the limit
func solve(solutions *int, limit int) {
	if *solutions >= limit {
		return
	}

	row, col, empty := findEmpty()
	if !empty {
		*solutions++
		// If it's the first solution, copy it to the solution board
		if *solutions == 1 {
			copySolution()
		}
		return
	}

	for num := 1; num <= 9; num++ {
		if isSafe(row, col, num) {
			board[row][col] = num
			solve(solutions, limit)
			board[row][col] = 0
		}
	}
}

// Find the next empty cell
func findEmpty() (int, int, bool) {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			if board[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

// Check if it's safe to place a number in a cell
func isSafe(row, col, num int) bool {
	// Check row and column
	for i := 0; i < SIZE; i++ {
		if board[row][i] == num || board[i][col] == num {
			return false
		}
	}

	// Check 3x3 subgrid
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

// Copy the current board to the solution board
func copySolution() {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			solution[i][j] = board[i][j]
		}
	}
}

// Print the solved Sudoku board
func printBoard(b [SIZE][SIZE]int) {
	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			fmt.Printf("%d", b[i][j])
			if j != SIZE-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
