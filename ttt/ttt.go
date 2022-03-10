package main

import "fmt"

// Make enum `Square` representing possible state of a board square (X, O, blank)
type Square int

const (
	Blank Square = iota
	X
	O
)

func (sq Square) token() string {
	switch sq {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return " "
	}
}

// `Board` is 2D array of `Square`
type Board [][]Square

// Print a board to console
func (b Board) print() {
	for _, row := range b {
		for _, sq := range row {
			fmt.Print(sq.token())
			fmt.Print("|")
		}
		fmt.Println()
		for range row {
			fmt.Print("--")
		}
		fmt.Println()
	}
}

// Who has won (blank -> no winner)
func (b Board) whoWon() Square {
	// Check rows/cols
	for i := 0; i < 3; i++ {
		// Check row `i`
		if b[i][0] == b[i][1] && b[i][1] == b[i][2] {
			return b[i][0]
		}
		// Check column `i`
		if b[0][i] == b[1][i] && b[1][i] == b[2][i] {
			return b[0][i]
		}
	}

	// Check diagonals
	if b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		return b[0][0]
	}
	if b[2][0] == b[1][1] && b[1][1] == b[0][2] {
		return b[2][0]
	}

	return Blank
}

func main() {
	// Initialise empty board
	board := Board{
		{X, Blank, Blank},
		{Blank, X, Blank},
		{Blank, Blank, X},
	}

	board.print()
	fmt.Println(board.whoWon().token())
}
