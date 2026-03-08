package main

import "fmt"

var shipTypes = []struct {
	name string
	size int
}{
	{"Carrier", 5},
	{"Battleship", 4},
	{"Cruiser", 3},
	{"Submarines", 3},
	{"Destroyer", 2},
}

type Board [boardSize][boardSize]string

type Position struct {
	row, col int
}

func printBoards(playerBoard, opponentBoardView *Board) {
	// Clear the screen
	fmt.Print("\033[H\033[2J\033[3J")
	fmt.Println("\n=== BATTLESHIP ===")
	fmt.Println()

	fmt.Println("  OPPONENT'S BOARD:")
	fmt.Println(headerRow)
	// Print opponent's board
	// Player does not see enemy's ship only hit and miss
	for i := range boardSize {
		fmt.Printf("%d ", i)
		for j := range boardSize {
			cell := opponentBoardView[i][j]
			switch cell {
			case ship:
				// Don't show ships
				fmt.Printf("%s ", hiddenShip)
			default:
				fmt.Printf("%s ", cell)
			}
		}
		fmt.Println()
	}

	fmt.Println("\n  YOUR BOARD:")
	fmt.Println(headerRow)

	// Print player's board - all info

	for i := range boardSize {
		fmt.Printf("%d ", i)
		for j := range boardSize {
			fmt.Printf("%s ", playerBoard[i][j])
		}
		fmt.Println()

	}
	fmt.Println()
}
