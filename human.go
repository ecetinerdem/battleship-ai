package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Ship struct {
	ShipName      string
	StartPosition Position
	EndPosition   Position
}

type HumanPlayer struct {
	board    Board
	ships    []Ship
	opponent *AIPlayer
}

func NewHumanPlayer() *HumanPlayer {
	p := &HumanPlayer{}
	for i := range boardSize {
		for j := range boardSize {
			p.board[i][j] = empty
		}
	}

	return p
}

func (p *HumanPlayer) TakeTurn(opponentBoard *Board) (Position, bool) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nEnter target position (e.g. A0)")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToUpper(input))

		if len(input) <= 2 {
			fmt.Println("Invalid format, use format like 'A0'")
			continue
		}

		if input[0] < 'A' || input[0] > 'j' {
			fmt.Println("Column must be between A and J")
			continue
		}

		col := int(input[0] - 'A')
		rowStr := input[1:]

		row, err := strconv.Atoi(rowStr)

		if err != nil || row < 0 || row >= boardSize {
			fmt.Println("Row must be between 0 and 9")
			continue
		}
	}
}

func (p *HumanPlayer) GetBoard() *Board {
	return &p.board
}

func (p *HumanPlayer) PlaceShips() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n=== SHIP PLACEMENT ==")
	fmt.Println("Place your ships on the board")
	fmt.Println("Format: A0 H (A0 starting position, H=horizontal or V=vertical)")
	fmt.Println("Positions are given as letter (A-J) for column and number (0-9) for row")
	fmt.Println("Press Enter to continue...")
	reader.ReadString('\n')

	for _, shipType := range shipTypes {
		for {
			// Display current board
			printBoards(&p.board, &Board{})
			fmt.Printf("\n Place your %s (length %d): ", shipType.name, shipType.size)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(strings.ToUpper(input))
			parts := strings.Fields(input)

			if len(parts) != 2 {
				fmt.Println("Invalid format! Use format like 'A0 H'")
				time.Sleep(2 * time.Second)
				continue
			}
			pos := parts[0]
			dir := parts[1]
			if len(pos) < 2 || (dir != "H" && dir != "V") {
				fmt.Println("Invalid input! Position must be like 'A0' and direction must be 'H' or 'V'")
				time.Sleep(2 * time.Second)
				continue
			}

			// Extract column which is a letter
			if pos[0] < 'A' || pos[0] > 'j' {
				fmt.Println("Column  must be between A-J!")
				time.Sleep(2 * time.Second)
				continue
			}

			col := int(pos[0] - 'A')

			// Extracting row which is a number
			rowString := pos[1:]
			row, err := strconv.Atoi(rowString)
			if err != nil || row < 0 || row >= boardSize {
				fmt.Println("Row must be between 0-9!")
				time.Sleep(2 * time.Second)
				continue
			}

			// Check if placement is valid
			valid := true
			positions := []Position{}

			for i := range shipType.size {
				var r, c int
				if dir == "H" {
					r, c = row, col+i // Horizantal placement column increases
				} else {
					r, c = row+i, col // Vertical placement row increases
				}

				// Check if ship would go off of board
				if r >= boardSize || col >= boardSize {
					valid = false
					fmt.Printf("Ship would go off of the board! (Attemted to place at position %c%d)\n", 'A'+c, r)
					time.Sleep(2 * time.Second)
					break
				}
				// Check if position overlaps with another ship

				if p.board[r][c] == ship {
					valid = false
					fmt.Printf("Ship is overlapping at position %c%d)\n", 'A'+c, r)
					time.Sleep(2 * time.Second)
					break
				}
				positions = append(positions, Position{r, c})
			}

			if valid {
				// Place the ship
				newShip := Ship{
					StartPosition: positions[0],
				}

				for _, pos := range positions {
					p.board[pos.row][pos.col] = ship
				}
				newShip.EndPosition = positions[len(positions)-1]
				newShip.ShipName = shipType.name
				p.ships = append(p.ships, newShip)
				break
			}
		}
	}

	// Show final placement
	printBoards(&p.board, &Board{})
	fmt.Println("\nAll ships placed! Press Enter the Start the game...")
	reader.ReadString('\n')

}
