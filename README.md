# Detailed Code Explanation: Building a Sudoku Solver in Go

In this section, we'll delve deeply into the Sudoku solver implemented in Go. We'll examine each part of the code line by line, explaining its purpose and how it contributes to the overall functionality of the program. Whether you're new to Go or looking to enhance your understanding of algorithms and programming concepts, this detailed walkthrough will help you grasp the intricacies of the Sudoku solver.

---

## Complete Code Overview

Before we begin the detailed explanation, here's the complete `main.go` code for reference:

```go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const SIZE = 9

// Sudoku board represented as a 2D array
var board [SIZE][SIZE]int

func main() {
	// Parse and validate input
	if !parseInput(os.Args) || !isValid() {
		fmt.Println("Error")
		return
	}

	// Solve the Sudoku and check for uniqueness
	solutions := 0
	solve(&solutions, 2)

	if solutions == 1 {
		printBoard()
	} else {
		fmt.Println("Error")
	}
}

func parseInput(args []string) bool {
	if len(args) != SIZE+1 {
		return false
	}

	for i := 0; i < SIZE; i++ {
		line := args[i+1]
		if len(line) != SIZE {
			return false
		}
		for j, ch := range line {
			if ch == '.' {
				board[i][j] = 0
			} else if ch >= '1' && ch <= '9' {
				board[i][j] = int(ch - '0')
			} else {
				return false
			}
		}
	}
	return true
}

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

func solve(solutions *int, limit int) bool {
	if *solutions >= limit {
		return true
	}

	row, col, empty := findEmpty()
	if !empty {
		*solutions++
		return false
	}

	for num := 1; num <= 9; num++ {
		if isSafe(row, col, num) {
			board[row][col] = num
			if solve(solutions, limit) {
				return true
			}
			board[row][col] = 0
		}
	}
	return false
}

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

func printBoard() {
	for i := 0; i < SIZE; i++ {
		var elems []string
		for j := 0; j < SIZE; j++ {
			elems = append(elems, strconv.Itoa(board[i][j]))
		}
		fmt.Println(strings.Join(elems, " "))
	}
}
```

Now, let's break down and explain each part of this code in detail.

---

## 1. Package Declaration and Imports

```go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)
```

### Explanation:

- **`package main`**
  
  - In Go, every executable program must have a `main` package. The `main` package tells the Go compiler that this package should compile as an executable program instead of a shared library.

- **`import` Statement**

  - The `import` block includes packages that the program depends on:
    - **`fmt`**: Provides functions for formatted I/O, such as printing to the console.
    - **`os`**: Offers a platform-independent interface to operating system functionality, including access to command-line arguments.
    - **`strconv`**: Contains functions for converting strings to other data types and vice versa.
    - **`strings`**: Provides functions for manipulating UTF-8 encoded strings.

---

## 2. Constants and Global Variables

```go
const SIZE = 9

// Sudoku board represented as a 2D array
var board [SIZE][SIZE]int
```

### Explanation:

- **`const SIZE = 9`**
  
  - Defines a constant named `SIZE` with a value of `9`. This constant represents the dimensions of the Sudoku board, which is a 9x9 grid.

- **`var board [SIZE][SIZE]int`**
  
  - Declares a global variable `board` as a two-dimensional array with dimensions `SIZE x SIZE` (i.e., 9x9).
  - Each element of the array is of type `int`, representing the numbers in the Sudoku grid. A value of `0` indicates an empty cell.

---

## 3. The `main` Function

```go
func main() {
	// Parse and validate input
	if !parseInput(os.Args) || !isValid() {
		fmt.Println("Error")
		return
	}

	// Solve the Sudoku and check for uniqueness
	solutions := 0
	solve(&solutions, 2)

	if solutions == 1 {
		printBoard()
	} else {
		fmt.Println("Error")
	}
}
```

### Explanation:

The `main` function is the entry point of the Go program. Here's what each part does:

1. **Parsing and Validating Input:**

   ```go
   if !parseInput(os.Args) || !isValid() {
	   fmt.Println("Error")
	   return
   }
   ```

   - **`parseInput(os.Args)`**
     - Calls the `parseInput` function, passing `os.Args` (the slice of command-line arguments) to it.
     - The function parses the input arguments and populates the `board` array.
     - Returns `true` if parsing is successful; `false` otherwise.

   - **`isValid()`**
     - After parsing, `isValid` checks whether the initial Sudoku board adheres to Sudoku rules (no duplicates in rows, columns, or subgrids).
     - Returns `true` if the board is valid; `false` otherwise.

   - **Error Handling:**
     - If either parsing fails or the board is invalid, the program prints `"Error"` and terminates early using `return`.

2. **Solving the Sudoku and Checking for Uniqueness:**

   ```go
   solutions := 0
   solve(&solutions, 2)
   ```

   - **`solutions := 0`**
     - Initializes a counter `solutions` to keep track of the number of solutions found.

   - **`solve(&solutions, 2)`**
     - Calls the `solve` function, passing the address of `solutions` and the limit `2`.
     - The limit `2` is used to check for uniqueness; if more than one solution is found, the function stops early.

3. **Output Based on Solutions Found:**

   ```go
   if solutions == 1 {
	   printBoard()
   } else {
	   fmt.Println("Error")
   }
   ```

   - **`if solutions == 1`**
     - Checks if exactly one solution was found.
     - If so, it calls `printBoard` to display the solved Sudoku.

   - **`else`**
     - If no solution or multiple solutions are found, it prints `"Error"` to indicate an invalid puzzle or non-unique solution.

---

## 4. Parsing Input: The `parseInput` Function

```go
func parseInput(args []string) bool {
	if len(args) != SIZE+1 {
		return false
	}

	for i := 0; i < SIZE; i++ {
		line := args[i+1]
		if len(line) != SIZE {
			return false
		}
		for j, ch := range line {
			if ch == '.' {
				board[i][j] = 0
			} else if ch >= '1' && ch <= '9' {
				board[i][j] = int(ch - '0')
			} else {
				return false
			}
		}
	}
	return true
}
```

### Explanation:

The `parseInput` function processes the command-line arguments to populate the Sudoku board. Here's a breakdown:

1. **Function Signature:**

   ```go
   func parseInput(args []string) bool
   ```

   - **`args []string`**
     - Receives a slice of strings representing the command-line arguments (`os.Args`).

   - **Returns `bool`**
     - Returns `true` if parsing is successful; `false` otherwise.

2. **Checking the Number of Arguments:**

   ```go
   if len(args) != SIZE+1 {
	   return false
   }
   ```

   - **`len(args) != SIZE+1`**
     - `os.Args` includes the program name as the first argument.
     - Therefore, for a 9x9 Sudoku, there should be `9` arguments plus the program name, totaling `SIZE+1` (`10`).

   - **Return `false`**
     - If the number of arguments is incorrect, parsing fails.

3. **Iterating Over Each Row:**

   ```go
   for i := 0; i < SIZE; i++ {
	   line := args[i+1]
	   if len(line) != SIZE {
		   return false
	   }
	   for j, ch := range line {
		   if ch == '.' {
			   board[i][j] = 0
		   } else if ch >= '1' && ch <= '9' {
			   board[i][j] = int(ch - '0')
		   } else {
			   return false
		   }
	   }
   }
   ```

   - **`for i := 0; i < SIZE; i++`**
     - Iterates over each row of the Sudoku board.

   - **`line := args[i+1]`**
     - Retrieves the `i`-th row from the command-line arguments.
     - `args[0]` is the program name, so rows start from `args[1]`.

   - **`if len(line) != SIZE`**
     - Ensures that each row contains exactly `9` characters.
     - Returns `false` if the row length is incorrect.

   - **Iterating Over Each Character in the Row:**

     ```go
     for j, ch := range line {
	     if ch == '.' {
		     board[i][j] = 0
	     } else if ch >= '1' && ch <= '9' {
		     board[i][j] = int(ch - '0')
	     } else {
		     return false
	     }
     }
     ```

     - **`for j, ch := range line`**
       - Iterates over each character `ch` in the current row `line`.
       - `j` represents the column index.

     - **`if ch == '.'`**
       - If the character is a dot (`.`), it signifies an empty cell.
       - **`board[i][j] = 0`**: Sets the corresponding cell to `0`.

     - **`else if ch >= '1' && ch <= '9'`**
       - Checks if the character is a digit between `1` and `9`.
       - **`board[i][j] = int(ch - '0')`**: Converts the character digit to its integer value and assigns it to the cell.

     - **`else`**
       - If the character is neither a dot nor a valid digit, parsing fails.
       - **`return false`**: Indicates invalid input.

4. **Successful Parsing:**

   ```go
   return true
   ```

   - If all rows and characters are valid, the function returns `true`, indicating successful parsing.

---

## 5. Validating the Initial Board: The `isValid` Function

```go
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
```

### Explanation:

The `isValid` function ensures that the initial Sudoku board doesn't violate Sudoku rules. It checks for duplicates in rows, columns, and 3x3 subgrids.

1. **Function Signature:**

   ```go
   func isValid() bool
   ```

   - **Returns `bool`**
     - Returns `true` if the board is valid; `false` otherwise.

2. **Checking Rows and Columns:**

   ```go
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
   ```

   - **`for i := 0; i < SIZE; i++`**
     - Iterates over each row and column index.

   - **`row := make([]bool, SIZE+1)` and `col := make([]bool, SIZE+1)`**
     - Creates two boolean slices to track the presence of numbers in the current row and column.
     - Index `0` is unused for simplicity, so indices `1-9` correspond to Sudoku numbers.

   - **`for j := 0; j < SIZE; j++`**
     - Iterates over each cell in the current row `i` and column `i`.

   - **Checking the Row:**

     ```go
     if board[i][j] != 0 {
	     if row[board[i][j]] {
		     return false
	     }
	     row[board[i][j]] = true
     }
     ```

     - **`if board[i][j] != 0`**
       - Only consider filled cells (non-zero).

     - **`if row[board[i][j]]`**
       - Checks if the current number has already been encountered in this row.
       - If `true`, a duplicate exists, violating Sudoku rules.

     - **`row[board[i][j]] = true`**
       - Marks the number as seen in the current row.

   - **Checking the Column:**

     ```go
     if board[j][i] != 0 {
	     if col[board[j][i]] {
		     return false
	     }
	     col[board[j][i]] = true
     }
     ```

     - **`if board[j][i] != 0`**
       - Only consider filled cells.

     - **`if col[board[j][i]]`**
       - Checks for duplicates in the current column.

     - **`col[board[j][i]] = true`**
       - Marks the number as seen in the current column.

   - **Outcome:**
     - If any duplicate is found in any row or column, the function returns `false`.

3. **Checking 3x3 Subgrids:**

   ```go
   for i := 0; i < SIZE; i += 3 {
	   for j := 0; j < SIZE; j += 3 {
		   if !isValidSubgrid(i, j) {
			   return false
		   }
	   }
   }
   ```

   - **`for i := 0; i < SIZE; i += 3` and `for j := 0; j < SIZE; j += 3`**
     - Iterates over the starting indices of each 3x3 subgrid.
     - The outer loop increments `i` by `3` to move to the next set of rows.
     - The inner loop increments `j` by `3` to move to the next set of columns.

   - **`if !isValidSubgrid(i, j)`**
     - Calls the `isValidSubgrid` function to check the validity of the subgrid starting at `(i, j)`.
     - If any subgrid is invalid (contains duplicates), the function returns `false`.

4. **Valid Board:**

   ```go
   return true
   ```

   - If all rows, columns, and subgrids are valid (no duplicates), the board is considered valid.

---

## 6. Validating a Subgrid: The `isValidSubgrid` Function

```go
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
```

### Explanation:

The `isValidSubgrid` function checks a specific 3x3 subgrid for duplicates.

1. **Function Signature:**

   ```go
   func isValidSubgrid(startRow, startCol int) bool
   ```

   - **Parameters:**
     - **`startRow`**: The starting row index of the subgrid.
     - **`startCol`**: The starting column index of the subgrid.

   - **Returns `bool`**
     - Returns `true` if the subgrid is valid; `false` otherwise.

2. **Tracking Seen Numbers:**

   ```go
   marks := make([]bool, SIZE+1)
   ```

   - Creates a boolean slice `marks` to track numbers encountered in the subgrid.
   - Index `0` is unused; indices `1-9` correspond to Sudoku numbers.

3. **Iterating Over the 3x3 Subgrid:**

   ```go
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
   ```

   - **`for i := 0; i < 3; i++` and `for j := 0; j < 3; j++`**
     - Iterates over the 3 rows and 3 columns of the subgrid.

   - **`val := board[startRow+i][startCol+j]`**
     - Retrieves the value at the current cell within the subgrid.

   - **`if val != 0`**
     - Only consider filled cells.

   - **`if marks[val]`**
     - Checks if the number `val` has already been encountered in this subgrid.
     - If `true`, a duplicate exists, violating Sudoku rules.

   - **`marks[val] = true`**
     - Marks the number as seen in the current subgrid.

4. **Valid Subgrid:**

   ```go
   return true
   ```

   - If no duplicates are found in the subgrid, the function returns `true`.

---

## 7. Solving the Sudoku: The `solve` Function

```go
func solve(solutions *int, limit int) bool {
	if *solutions >= limit {
		return true
	}

	row, col, empty := findEmpty()
	if !empty {
		*solutions++
		return false
	}

	for num := 1; num <= 9; num++ {
		if isSafe(row, col, num) {
			board[row][col] = num
			if solve(solutions, limit) {
				return true
			}
			board[row][col] = 0
		}
	}
	return false
}
```

### Explanation:

The `solve` function employs a backtracking algorithm to find solutions to the Sudoku puzzle. It also checks for the uniqueness of the solution by limiting the number of solutions to `2`.

1. **Function Signature:**

   ```go
   func solve(solutions *int, limit int) bool
   ```

   - **Parameters:**
     - **`solutions *int`**: A pointer to an integer that counts the number of solutions found.
     - **`limit int`**: The maximum number of solutions to search for before stopping.

   - **Returns `bool`**
     - Returns `true` if the solution limit is reached; `false` otherwise.

2. **Early Termination if Limit Reached:**

   ```go
   if *solutions >= limit {
	   return true
   }
   ```

   - Checks if the number of solutions found has reached or exceeded the `limit`.
   - If so, returns `true` to signal that further searching can be stopped.

3. **Finding the Next Empty Cell:**

   ```go
   row, col, empty := findEmpty()
   if !empty {
	   *solutions++
	   return false
   }
   ```

   - **`findEmpty()`**
     - Calls the `findEmpty` function to locate the next empty cell in the Sudoku board.
     - Returns the row and column indices of the empty cell and a boolean indicating whether an empty cell was found.

   - **`if !empty`**
     - If no empty cells are found (`empty` is `false`), it means the board is completely filled.
     - **`*solutions++`**: Increments the `solutions` counter.
     - **`return false`**: Returns `false` to continue searching for more solutions until the limit is reached.

4. **Trying Possible Numbers:**

   ```go
   for num := 1; num <= 9; num++ {
	   if isSafe(row, col, num) {
		   board[row][col] = num
		   if solve(solutions, limit) {
			   return true
		   }
		   board[row][col] = 0
	   }
   }
   ```

   - **`for num := 1; num <= 9; num++`**
     - Iterates through numbers `1` to `9`, attempting to place each in the current empty cell.

   - **`if isSafe(row, col, num)`**
     - Calls the `isSafe` function to check if placing `num` in the cell `(row, col)` is valid according to Sudoku rules.

   - **Placing the Number:**

     ```go
     board[row][col] = num
     ```

     - Temporarily assigns `num` to the cell `(row, col)`.

   - **Recursive Call:**

     ```go
     if solve(solutions, limit) {
	     return true
     }
     ```

     - Recursively calls `solve` to attempt to solve the rest of the board with the current number placed.
     - If the recursive call returns `true`, it indicates that the solution limit has been reached, and the function returns `true` to stop further searching.

   - **Backtracking:**

     ```go
     board[row][col] = 0
     ```

     - If placing `num` doesn't lead to a valid solution, the function backtracks by resetting the cell to `0` (empty) and tries the next number.

5. **No Valid Number Found:**

   ```go
   return false
   ```

   - If no valid number can be placed in the current empty cell, the function returns `false` to trigger backtracking.

---

## 8. Finding the Next Empty Cell: The `findEmpty` Function

```go
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
```

### Explanation:

The `findEmpty` function searches for the next empty cell (marked as `0`) in the Sudoku board.

1. **Function Signature:**

   ```go
   func findEmpty() (int, int, bool)
   ```

   - **Returns:**
     - **`int`**: Row index of the empty cell.
     - **`int`**: Column index of the empty cell.
     - **`bool`**: `true` if an empty cell is found; `false` otherwise.

2. **Iterating Over the Board:**

   ```go
   for i := 0; i < SIZE; i++ {
	   for j := 0; j < SIZE; j++ {
		   if board[i][j] == 0 {
			   return i, j, true
		   }
	   }
   }
   ```

   - **`for i := 0; i < SIZE; i++`**
     - Iterates over each row.

   - **`for j := 0; j < SIZE; j++`**
     - Iterates over each column within the current row.

   - **`if board[i][j] == 0`**
     - Checks if the current cell is empty.

   - **`return i, j, true`**
     - If an empty cell is found, returns its row and column indices along with `true`.

3. **No Empty Cell Found:**

   ```go
   return -1, -1, false
   ```

   - If the entire board is filled (no empty cells), returns `-1` for both row and column indices and `false` to indicate that no empty cell was found.

---

## 9. Checking if a Move is Safe: The `isSafe` Function

```go
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
```

### Explanation:

The `isSafe` function determines whether placing a specific number in a given cell violates Sudoku rules.

1. **Function Signature:**

   ```go
   func isSafe(row, col, num int) bool
   ```

   - **Parameters:**
     - **`row`**: Row index where the number is to be placed.
     - **`col`**: Column index where the number is to be placed.
     - **`num`**: The number to be placed in the cell.

   - **Returns `bool`**
     - Returns `true` if placing `num` in `(row, col)` is safe; `false` otherwise.

2. **Checking the Row and Column:**

   ```go
   for i := 0; i < SIZE; i++ {
	   if board[row][i] == num || board[i][col] == num {
		   return false
	   }
   }
   ```

   - **`for i := 0; i < SIZE; i++`**
     - Iterates over each cell in the specified row and column.

   - **`if board[row][i] == num || board[i][col] == num`**
     - Checks if `num` already exists in the same row or column.
     - If found, placing `num` would violate Sudoku rules.

   - **`return false`**
     - Returns `false` immediately if a duplicate is found.

3. **Checking the 3x3 Subgrid:**

   ```go
   startRow := (row / 3) * 3
   startCol := (col / 3) * 3
   for i := 0; i < 3; i++ {
	   for j := 0; j < 3; j++ {
		   if board[startRow+i][startCol+j] == num {
			   return false
		   }
	   }
   }
   ```

   - **`startRow := (row / 3) * 3` and `startCol := (col / 3) * 3`**
     - Determines the starting indices of the 3x3 subgrid that contains the cell `(row, col)`.
     - For example, if `row` is `5`, `(5 / 3) * 3 = 3`, so the subgrid starts at row `3`.

   - **Nested Loops:**

     ```go
     for i := 0; i < 3; i++ {
	     for j := 0; j < 3; j++ {
		     if board[startRow+i][startCol+j] == num {
			     return false
		     }
	     }
     }
     ```

     - Iterates over the 3 rows and 3 columns of the subgrid.
     - **`if board[startRow+i][startCol+j] == num`**
       - Checks if `num` already exists in the subgrid.
       - If found, placing `num` would violate Sudoku rules.

     - **`return false`**
       - Returns `false` immediately if a duplicate is found.

4. **Safe to Place the Number:**

   ```go
   return true
   ```

   - If `num` is not found in the same row, column, or subgrid, it's safe to place it.
   - Returns `true` to indicate that placing `num` in `(row, col)` does not violate any Sudoku rules.

---

## 10. Printing the Solved Board: The `printBoard` Function

```go
func printBoard() {
	for i := 0; i < SIZE; i++ {
		var elems []string
		for j := 0; j < SIZE; j++ {
			elems = append(elems, strconv.Itoa(board[i][j]))
		}
		fmt.Println(strings.Join(elems, " "))
	}
}
```

### Explanation:

The `printBoard` function displays the solved Sudoku board in a readable format, with numbers separated by spaces.

1. **Function Signature:**

   ```go
   func printBoard()
   ```

   - **No Parameters or Return Values**
     - The function operates directly on the global `board` variable and prints the result.

2. **Iterating Over Each Row:**

   ```go
   for i := 0; i < SIZE; i++ {
	   var elems []string
	   for j := 0; j < SIZE; j++ {
		   elems = append(elems, strconv.Itoa(board[i][j]))
	   }
	   fmt.Println(strings.Join(elems, " "))
   }
   ```

   - **`for i := 0; i < SIZE; i++`**
     - Iterates over each row of the Sudoku board.

   - **`var elems []string`**
     - Initializes a slice of strings to hold the numbers of the current row as strings.

   - **`for j := 0; j < SIZE; j++`**
     - Iterates over each column in the current row.

   - **`elems = append(elems, strconv.Itoa(board[i][j]))`**
     - Converts the integer value in `board[i][j]` to a string using `strconv.Itoa`.
     - Appends the string to the `elems` slice.

   - **`fmt.Println(strings.Join(elems, " "))`**
     - Joins the elements of the `elems` slice into a single string, separated by spaces.
     - Prints the joined string, resulting in a neatly formatted row of numbers.

3. **Result:**

   - The function prints each row of the Sudoku board on a separate line, with numbers separated by spaces, matching the expected output format.

---

## Summary of Functionality

Here's a high-level overview of how the entire program operates:

1. **Input Parsing:**
   - The program expects exactly nine command-line arguments (excluding the program name), each representing a row of the Sudoku puzzle.
   - Each row should consist of nine characters: digits (`1-9`) for filled cells and dots (`.`) for empty cells.

2. **Validation:**
   - After parsing, the program validates the initial board to ensure there are no duplicates in any row, column, or 3x3 subgrid.

3. **Solving:**
   - Uses a recursive backtracking algorithm to fill in empty cells.
   - The solver attempts to place numbers `1-9` in empty cells, ensuring each placement is safe.
   - Counts the number of valid solutions found to ensure uniqueness.

4. **Output:**
   - If exactly one solution is found, the program prints the completed Sudoku board.
   - If no solution or multiple solutions are found, it prints `"Error"` to indicate an invalid or ambiguous puzzle.

---

## Additional Notes

- **Global Variable Usage:**
  - The `board` is defined as a global variable for ease of access across different functions. While this approach simplifies the code for small projects, in larger applications, it's advisable to encapsulate data within structures or pass them explicitly to functions to enhance modularity and maintainability.

- **Error Handling:**
  - The program employs simple error handling by printing `"Error"` for any invalid input or unsolvable puzzles. For more informative error messages, consider specifying the exact reason for the failure (e.g., "Duplicate number in row 3").

- **Performance Considerations:**
  - The backtracking algorithm is efficient for standard 9x9 Sudoku puzzles. However, for larger puzzles or puzzles with many empty cells, performance optimizations (like implementing heuristics) can significantly reduce solving time.

- **Extensibility:**
  - The current implementation is tailored for 9x9 Sudoku puzzles. To adapt it for different sizes (e.g., 16x16), you'd need to adjust the `SIZE` constant, validation logic, and input parsing accordingly.

---
