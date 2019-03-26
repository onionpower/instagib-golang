package main

import (
	"errors"
	"fmt"
)

const Rook = 'R'
const Pawn = 'p'
const Empty = '.'
const FieldLen = 8

func numRookCaptures(board [][]byte) int {
	x, y := findRook(board)
	captures := 0
	for h := x - 1; h >= 0; h-- {
		if board[y][h] == Pawn {
			captures++
			break
		}
		if board[y][h] != Empty {
			break
		}
	}
	for h := x + 1; h < FieldLen; h++ {
		if board[y][h] == Pawn {
			captures++
			break
		}
		if board[y][h] != Empty {
			break
		}
	}
	for h := y - 1; h >= 0; h-- {
		if board[y][h] == Pawn {
			captures++
			break
		}
		if board[y][h] != Empty {
			break
		}
	}
	for h := y + 1; h < FieldLen; h++ {
		if board[y][h] == Pawn {
			captures++
			break
		}
		if board[y][h] != Empty {
			break
		}
	}
	return captures
}

func findRook(bytes [][]byte) (int, int) {
	for x, row := range bytes {
		for y, cell := range row {
			if cell == Rook {
				return y, x
			}
		}
	}

	panic(errors.New("it's a lie!"))
}

func main() {
	a := [][]string{
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", "R", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."}}
	printSln(a)
}

func toBytes(s [][]string) [][]byte {
	res := make([][]byte, 0, 8)
	for _, r := range s {
		resRow := make([]byte, 0, 8)
		for _, s := range r {
			resRow = append(resRow, byte(s[0]))
		}
		res = append(res, resRow)
	}
	return res
}

func printSln(board [][]string) {
	a := toBytes(board)
	fmt.Println(board)
	fmt.Printf("captures: %v\n", numRookCaptures(a))
}
