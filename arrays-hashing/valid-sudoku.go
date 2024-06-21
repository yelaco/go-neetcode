package arrays_hashing

import "fmt"

func isValidSudoku(board [][]byte) bool {
	rows := [9]map[byte]struct{}{}
	cols := [9]map[byte]struct{}{}
	squares := [9]map[byte]struct{}{}

	for i := 0; i < 9; i++ {
		rows[i] = map[byte]struct{}{}
		cols[i] = map[byte]struct{}{}
		squares[i] = map[byte]struct{}{}
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}

			if _, ok := rows[r][board[r][c]]; ok {
				return false
			}

			if _, ok := cols[c][board[r][c]]; ok {
				return false
			}

			squareIdx := (r/3)*3 + (c / 3)
			if _, ok := squares[squareIdx][board[r][c]]; ok {
				return false
			}

			rows[r][board[r][c]] = struct{}{}
			cols[c][board[r][c]] = struct{}{}
			squares[squareIdx][board[r][c]] = struct{}{}
		}
	}

	return true
}

func TestValidSudoku() {
	testCases := [][][]byte{
		{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'}},
	}

	for _, testCase := range testCases {
		fmt.Printf("Input: board =\n%v\nOutput: %v\n", testCase, isValidSudoku(testCase))
		fmt.Println()
	}
}
