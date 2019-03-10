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
	// find 'R'
	// draw a cross
	// check if 'p' is the closest figure. Skip
	// all dots, but stop on everything else
	x, y := findRook(board)
	captures := 0
	captures += findCapturesInTheLeft(board, x, y)
	captures += findCapturesInTheRight(board, x, y)
	captures += findCapturesBelow(board, x, y)
	captures += findCapturesAbove(board, x, y)
	fmt.Printf("The Rook is on %v, %v\n", x, y)

	return captures
}

func findCapturesInTheLeft(board [][]byte, x int, y int) int {
	for h := x - 1; h >= 0; h-- {
		stop, captures := capture(board, h, y)
		if stop {
			return captures
		}
	}

	return 0
}

func findCapturesBelow(board [][]byte, x int, y int) int {
	for h := y - 1; h >= 0; h-- {
		stop, captures := capture(board, x, h)
		if stop {
			return captures
		}
	}

	return 0
}

func findCapturesAbove(board [][]byte, x int, y int) int {
	for h := y + 1; h < FieldLen; h++ {
		stop, captures := capture(board, x, h)
		if stop {
			return captures
		}
	}

	return 0
}

func findCapturesInTheRight(board [][]byte, x int, y int) int {
	for h := x + 1; h < FieldLen; h++ {
		stop, captures := capture(board, h, y)
		if stop {
			return captures
		}
	}

	return 0
}

func capture(board [][]byte, x int, y int) (bool, int) {
	it := board[y][x]
	if it == Empty {
		return false, 0
	}
	if it == Pawn {
		return true, 1
	}
	return true, 0
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
		{".", ".", ".", "p", ".", ".", ".", "."},
		{".", "p", ".", "R", ".", ".", ".", "p"},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", "p", ".", ".", ".", "."},
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
