package main

import (
	"fmt"
	"strings"
)

type Puzzle struct {
	Grid [9][9]int
}

func (b *Puzzle) Solve(forward bool) bool {
	var row, col int
	if forward {
		row, col = b.nextEmptyForward()
	} else {
		row, col = b.nextEmptyReverse()
	}

	if row < 0 {
		return true
	}
	for i := 1; i <= 9; i++ {
		if b.IsValid(i, row, col) {
			b.Grid[row][col] = i
			if b.Solve(forward) {
				return true
			}
			b.Grid[row][col] = 0
		}
	}
	return false
}

func (b *Puzzle) nextEmptyForward() (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b.Grid[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func (b *Puzzle) nextEmptyReverse() (int, int) {
	for i := 8; i >= 0; i-- {
		for j := 8; j >= 0; j-- {
			if b.Grid[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

func (b *Puzzle) IsValid(val, row, col int) bool {
	for i := 0; i < 9; i++ {
		if b.Grid[row][i] == val && col != i {
			return false
		}
		if b.Grid[i][col] == val && row != i {
			return false
		}
	}

	boxRow := row / 3
	boxCol := col / 3
	for i := boxRow * 3; i < boxRow*3+3; i++ {
		for j := boxCol * 3; j < boxCol*3+3; j++ {
			if b.Grid[i][j] == val && (i != row || j != col) {
				return false
			}
		}
	}
	return true
}

func (b Puzzle) String() string {
	var builder strings.Builder
	builder.Grow(172)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Fprintf(&builder, "%d ", b.Grid[i][j])
		}
		fmt.Fprint(&builder, "\n")
	}
	return builder.String()
}

func main() {
	puzzle := Puzzle{
		Grid: [9][9]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 3, 0, 8, 5},
			{0, 0, 1, 0, 2, 0, 0, 0, 0},
			{0, 0, 0, 5, 0, 7, 0, 0, 0},
			{0, 0, 4, 0, 0, 0, 1, 0, 0},
			{0, 9, 0, 0, 0, 0, 0, 0, 0},
			{5, 0, 0, 0, 0, 0, 0, 7, 3},
			{0, 0, 2, 0, 1, 0, 0, 0, 0},
			{0, 0, 0, 0, 4, 0, 0, 0, 9},
		},
	}
	puzzle2 := puzzle

	results := make(chan Puzzle)
	go func() {
		puzzle.Solve(true)
		results <- puzzle
	}()
	go func() {
		puzzle2.Solve(false)
		results <- puzzle2
	}()
	fmt.Println(<-results)
}
