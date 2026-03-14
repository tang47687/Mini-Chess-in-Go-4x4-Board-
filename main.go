package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Pieces: White (Uppercase), Black (Lowercase)
// K/k = King, P/p = Pawn
type Board [4][4]string

var (
	gameBoard Board
	turn      = "White"
)

func initBoard() {
	// Initialize Black pieces
	gameBoard[0] = [4]string{"k", "p", "p", "p"}
	// Initialize White pieces
	gameBoard[3] = [4]string{"K", "P", "P", "P"}
	// Fill empty squares
	for i := 1; i < 3; i++ {
		for j := 0; j < 4; j++ {
			gameBoard[i][j] = "."
		}
	}
}

func displayBoard() {
	fmt.Println("\n    0   1   2   3") // Column indices
	fmt.Println("  -----------------")
	for i, row := range gameBoard {
		fmt.Printf("%d | ", i) // Row index
		for _, cell := range row {
			fmt.Printf("%s   ", cell)
		}
		fmt.Println("|")
	}
	fmt.Println("  -----------------")
	fmt.Printf("Current Turn: %s\n", turn)
}

func isValidMove(r1, c1, r2, c2 int) (bool, string) {
	// Out of bounds check
	if r1 < 0 || r1 > 3 || c1 < 0 || c1 > 3 || r2 < 0 || r2 > 3 || c2 < 0 || c2 > 3 {
		return false, "Coordinates out of bounds!"
	}

	piece := gameBoard[r1][c1]
	target := gameBoard[r2][c2]

	// Check if selecting an empty square
	if piece == "." {
		return false, "No piece at the starting position!"
	}

	// Turn validation
	isWhitePiece := strings.ToUpper(piece) == piece
	if (turn == "White" && !isWhitePiece) || (turn == "Black" && isWhitePiece) {
		return false, "It is not your turn!"
	}

	// Prevent friendly fire
	if target != "." {
		isTargetWhite := strings.ToUpper(target) == target
		if isWhitePiece == isTargetWhite {
			return false, "You cannot capture your own piece!"
		}
	}

	// Simple movement logic: 1 square in any direction (King-like)
	dr := r2 - r1
	dc := c2 - c1
	if dr < -1 || dr > 1 || dc < -1 || dc > 1 {
		return false, "Invalid move! Pieces move only 1 square."
	}

	return true, ""
}

func main() {
	initBoard()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Go Mini-Chess!")
	fmt.Println("Format: row1 col1 row2 col2 (e.g., 3 1 2 1)")

	for {
		displayBoard()
		fmt.Print("Enter your move: ")
		
		if !scanner.Scan() {
			break
		}
		
		var r1, c1, r2, c2 int
		_, err := fmt.Sscanf(scanner.Text(), "%d %d %d %d", &r1, &c1, &r2, &c2)
		
		if err != nil {
			fmt.Println("Invalid input. Use: start_row start_col end_row end_col")
			continue
		}

		valid, msg := isValidMove(r1, c1, r2, c2)
		if !valid {
			fmt.Printf("Error: %s\n", msg)
			continue
		}

		// Execute move
		gameBoard[r2][c2] = gameBoard[r1][c1]
		gameBoard[r1][c1] = "."

		// Check for King capture (Win condition)
		// (Simplified logic: check if 'k' or 'K' was replaced)
		
		// Switch turn
		if turn == "White" {
			turn = "Black"
		} else {
			turn = "White"
		}
	}
}
