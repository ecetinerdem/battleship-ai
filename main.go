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
	ai := NewAIPlayer()

	human.opponent = ai
	ai.opponent = human

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
	ai.PlaceShips()
	// Human places ship
	human.PlaceShips()
	// Main game loop
	gameOver := false
	playerTurn := true

	for !gameOver {
		// Display the boards
		printBoards(human.GetBoard(), ai.GetBoard())

		// Players take turn
		if playerTurn {
			fmt.Println("\n=== YOUR TURN ===")
			// Let player take turn
			_, _ = human.TakeTurn(ai.GetBoard())

			// Check win condition
			if checkWinCondititon(ai.GetBoard()) {
				gameOver = true
				printBoards(human.GetBoard(), ai.GetBoard())
				fmt.Println("\nYou Won! You sank all enemy ships")
			}
		} else {

			// fmt.Println("Heat Map:")
			// for i := range boardSize {
			// 	for j := range boardSize {
			// 		fmt.Printf("%3d  ", ai.heatMap[i][j])
			// 	}
			// 	fmt.Println()
			// }

			fmt.Println("\n=== AI's TURN ===")
			// Let AI take turn
			_, _ = ai.TakeTurn(human.GetBoard())

			// fmt.Println("Press enter to continue...")
			// reader.ReadString('\n')

			// Check win condition
			if checkWinCondititon(human.GetBoard()) {
				gameOver = true
				printBoards(human.GetBoard(), ai.GetBoard())
				fmt.Println("\nYou Lost! All your ships sunk")
			}
		}
		// Switch turns
		playerTurn = !playerTurn

	}

	fmt.Println("\nThanks for playing Battleship! Press enter to exit...")
	reader.ReadString('\n')
}
