package _036_suduko

import "fmt"

func isValidSudukoRow(board [][]byte, row int) bool {
	var flags [10]bool
	for col := 0; col < 9; col++ {
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if flags[n] {
			fmt.Println("Invalid row: ", row)
			return false
		}
		flags[n] = true
	}
	if isEmpty(flags) {
		return false
	}
	return true
}

func isValidSudukoCol(board [][]byte, col int) bool {
	var flags[10] bool
	for row := 0; row < 9; row++ {
		n := convertToNumber(board[row][col])
		if n < 0 {
			continue
		}
		if flags[n] {
			fmt.Println("Invalid col: ", col)
			return false
		}
		flags[n] = true
	}
	return true
}

func isValidSudukoPod(board [][]byte, pod int) bool {
	drow, dcol := pod / 3, pod % 3
	subBoard := board[drow * 3:drow*4-1][dcol*3:dcol*4-1]
	var flags [10]bool
	var i int
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			i++
			n := convertToNumber(subBoard[row][col])
			if n < 0 {
				continue
			}
			if flags[i] {
				fmt.Println("Invalid pod: ", pod)
				return false
			}
			flags[i] = true
		}
	}
	if isEmpty(flags) {
		return false
	}
	return true
}

// every unit(line,column, subBox) must have at least one number between 1-9.
func isEmpty(flags [10]bool) bool {
	var hasTrue bool
	for _, flag := range flags {
		hasTrue = hasTrue && flag
	}
	if !hasTrue {
		return true
	}
	return false
}

func convertToNumber(b byte) int {
	if b == '.' {
		return -1
	}
	return int(b - '0')
}

func isValidSuduko(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		if !isValidSudukoRow(board, row) {
			return false
		}
	}
	for col := 0; col < 9; col++ {
		if !isValidSudukoCol(board, col) {
			return false
		}
	}
	for pod := 0; pod < 9; pod++ {
		if !isValidSudukoPod(board, pod) {
			return false
		}
	}
	return true
}
