package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	boardSize    = 10
	empty        = "."
	ship         = "O"
	hit          = "X"
	miss         = "~"
	hiddenShip   = "."
	headerRow    = "  A B C D E F G H I J"
	headerColumn = "0123456789"
)

func main() {
	// Create players
	human := NewHumanPlayer()

	// Welcome Message
	fmt.Println("\n=== WELCOME TO BATTLESHIP ===")
	fmt.Println("Legend:")
	fmt.Printf("   %s - Empty water\n", empty)
	fmt.Printf("   %s - Your ship\n", ship)
	fmt.Printf("   %s - Hit\n", hit)
	fmt.Printf("   %s - Miss\n", miss)
	fmt.Println("\nPress enter to start the game...")

	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	// Place the ships
	// AI place ships (player does not know)
	// Human places ship

	// Main game loop
	gameOver := false
	//playerTurn := true

	for !gameOver {
		// Display the boards

		// Players take turn

		// Switch turns

		// Check win condition
	}

	fmt.Println("\nThanks for playing Battleship! Press enter tı exit...")
	reader.ReadString('\n')
}
