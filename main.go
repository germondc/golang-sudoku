package main

import (
	"fmt"
	"time"
)

func getBoard(index int) [9][9]int {
	boards := [][9][9]int{{
		{8, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 3, 6, 0, 0, 0, 0, 0},
		{0, 7, 0, 0, 9, 0, 2, 0, 0},
		{0, 5, 0, 0, 0, 7, 0, 0, 0},
		{0, 0, 0, 0, 4, 5, 7, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 3, 0},
		{0, 0, 1, 0, 0, 0, 0, 6, 8},
		{0, 0, 8, 5, 0, 0, 0, 1, 0},
		{0, 9, 0, 0, 0, 0, 4, 0, 0},
	}, {
		{0, 0, 5, 3, 0, 0, 0, 0, 0},
		{8, 0, 0, 0, 0, 0, 0, 2, 0},
		{0, 7, 0, 0, 1, 0, 5, 0, 0},
		{4, 0, 0, 0, 0, 5, 3, 0, 0},
		{0, 1, 0, 0, 7, 0, 0, 0, 6},
		{0, 0, 3, 2, 0, 0, 0, 8, 0},
		{0, 6, 0, 5, 0, 0, 0, 0, 9},
		{0, 0, 4, 0, 0, 0, 0, 3, 0},
		{0, 0, 0, 0, 0, 9, 7, 0, 0},
	}, {
		{1, 0, 0, 0, 0, 7, 0, 9, 0},
		{0, 3, 0, 0, 2, 0, 0, 0, 8},
		{0, 0, 9, 6, 0, 0, 5, 0, 0},
		{0, 0, 5, 3, 0, 0, 9, 0, 0},
		{0, 1, 0, 0, 8, 0, 0, 0, 2},
		{6, 0, 0, 0, 0, 4, 0, 0, 0},
		{3, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 4, 0, 0, 0, 0, 0, 0, 7},
		{0, 0, 7, 0, 0, 0, 3, 0, 0},
	}}

	return boards[index]
}

func testRow(data *[9][9]int, row int, col int) bool {
	for i := 0; i < 9; i++ {
		if i != col && data[row][i] == data[row][col] {
			return false
		}
	}
	return true
}

func testCol(data *[9][9]int, row int, col int) bool {
	for i := 0; i < 9; i++ {
		if i != row && data[i][col] == data[row][col] {
			return false
		}
	}
	return true
}

func testBlock(data *[9][9]int, row int, col int) bool {
	rowIndex := (row / 3) * 3
	colIndex := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (rowIndex+i) != row && (colIndex+j) != col && data[rowIndex+i][colIndex+j] == data[row][col] {
				return false
			}
		}
	}
	return true
}

func solve(board *[9][9]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == 0 {
				for val := 1; val <= 9; val++ {
					board[row][col] = val
					if testRow(board, row, col) && testCol(board, row, col) && testBlock(board, row, col) && solve(board) {
						return true
					}
					board[row][col] = 0
				}
				return false
			}
		}
	}
	return true
}

func printBoard(board [9][9]int) {
	for row := 0; row < 9; row++ {
		if row%3 == 0 {
			fmt.Println("+-------+-------+-------+")
		}
		for col := 0; col < 9; col++ {
			if col%3 == 0 {
				fmt.Print("| ")
			}
			fmt.Printf("%d ", board[row][col])
		}
		fmt.Println("|")
	}
	fmt.Println("+-------+-------+-------+")
}

func main() {
	board := getBoard(0)
	start := time.Now()
	result := solve(&board)

	elapsed := time.Since(start)

	if result {
		printBoard(board)
	}

	fmt.Printf("elapsed = %s\n", elapsed)
}
