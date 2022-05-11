/**
* 36. Valid Sudoku (medium)
**/

package main

import (
	"strconv"
)

// hash set with bits
func isValidSudoku(board [][]byte) bool {
	n := len(board)
	subs := make([]int, n+1)
	rows := make([]int, n+1)
	cols := make([]int, n+1)

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if board[r][c] == '.' {
				continue
			}

			// num := int(board[r][c])
			num, _ := strconv.Atoi(string(board[r][c]))
			s := r/3*3 + c/3
			checkSub := subs[s]&(1<<num) != 0
			checkRow := rows[r]&(1<<num) != 0
			checkCol := cols[c]&(1<<num) != 0
			if checkSub || checkRow || checkCol {
				return false
			}

			subs[s] |= 1 << num
			rows[r] |= 1 << num
			cols[c] |= 1 << num
		}
	}

	return true
}

// hash set
// func isValidSudoku2(board [][]byte) bool {
//     //
// }
