package main

import (
	"fmt"
)

func valid_table(table []int, row int, col int) bool {
	for row_i := 0; row_i < row; row_i++ {
		col_i := table[row_i]
		delta := row - row_i
		if col == col_i || col == col_i+delta || col == col_i-delta {
			return false
		}
	}
	return true
}

func n_queens(table []int, row int) {
	n := len(table)
	if row == n {
		draw_table(table)
	} else {
		for col := 0; col < n; col++ {
			if valid_table(table, row, col) {
				table[row] = col
				n_queens(table, row+1)
			}
		}
	}
}

func draw_table(table []int) {
	n := len(table)
	for row := 0; row < n; row++ {
		row_string := ""
		for col := 0; col < n; col++ {
			if table[row] == col {
				row_string += "Q"
			} else {
				row_string += "*"
			}
		}
		fmt.Println(row_string)
	}
	fmt.Println()
}

func main() {
	n := 5
	table := make([]int, n)
	n_queens(table, 0)
}
