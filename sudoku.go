package main

import (
	"os"

	"github.com/01-edu/z01"
)

var puzzle [9][9]rune

func BoardValidity(guess rune, row, col int) bool {
	for j := 0; j < 9; j++ {
		if puzzle[row][j] == guess {
			return false
		}
	}
	for i := 0; i < 9; i++ {
		if puzzle[i][col] == guess {
			return false
		}
	}
	row_start := (row / 3) * 3
	col_start := (col / 3) * 3
	for i := row_start; i < row_start+3; i++ {
		for j := col_start; j < col_start+3; j++ {
			if puzzle[i][j] == guess {
				return false
			}
		}
	}
	return true
}

func NextEmpty() (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if puzzle[i][j] == '.' {
				return i, j
			}
		}
	}
	return -1, -1
}

func SudokuSolver() bool {
	row, col := NextEmpty()
	if row == -1 {
		return true
	}
	for i := 1; i < 10; i++ {
		if BoardValidity(rune(i+48), row, col) == true {
			puzzle[row][col] = rune(i + 48)
			if SudokuSolver() == true {
				return true
			}
		}
		puzzle[row][col] = '.'
	}
	return false
}

func ReportError() {
	error := "Error"
	for i := 0; i < len(error); i++ {
		z01.PrintRune(rune(error[i]))
	}
	z01.PrintRune('\n')
}

func main() {
	args := os.Args[1:]
	if len(args) != 9 {
		ReportError()
		return
	}
	for i := 0; i < 9; i++ {
		if len(args[i]) != 9 {
			ReportError()
			return
		}
		for j := 0; j < 9; j++ {
			if args[i][j] > '0' && args[i][j] <= '9' {
				if BoardValidity(rune(args[i][j]), i, j) == true {
					puzzle[i][j] = rune(args[i][j])
				} else {
					ReportError()
					return
				}
			} else {
				if args[i][j] != '.' {
					ReportError()
					return
				}
				puzzle[i][j] = '.'
			}
		}
	}
	if SudokuSolver() == true {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				z01.PrintRune(puzzle[i][j])
				if j != 8 {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	} else {
		ReportError()
		return
	}
}
