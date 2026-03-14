package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func checkWinCondititon(board *Board) bool {
	for i := range boardSize {
		for j := range boardSize {
			if board[i][j] == ship {
				return false
			}
		}
	}
	return true
}

func isShipSunk(board *Board, row, col int, opponentShips []Ship) (bool, string) {
	// Find the ship that was hit
	var hitShip *Ship

	for i := range opponentShips {
		ship := &opponentShips[i]
		// Check if hit coordinates are within this ships boundaries
		if (ship.StartPosition.row == ship.EndPosition.row &&
			row == ship.StartPosition.row &&
			col >= ship.StartPosition.col &&
			col <= ship.EndPosition.col) ||
			(ship.StartPosition.col == ship.EndPosition.col &&
				col == ship.StartPosition.col &&
				row >= ship.StartPosition.row &&
				row <= ship.EndPosition.row) {
			hitShip = ship
			break
		}
	}
	// If no ship was found at this hit location return false
	if hitShip == nil {
		fmt.Println("No ship was found at this location")
		return false, ""
	}
	// Check if all parts of the ship are marked as hit
	hits := 0
	shipLength := 0

	if hitShip.StartPosition.row == hitShip.EndPosition.row {
		//Horizontal ship
		shipLength = hitShip.EndPosition.col - hitShip.StartPosition.col + 1
		for c := hitShip.StartPosition.col; c <= hitShip.EndPosition.col; c++ {
			if board[hitShip.StartPosition.row][c] == hit {
				hits++
			}
		}
	} else {
		// Vertical ship
		shipLength = hitShip.EndPosition.row - hitShip.StartPosition.row + 1
		for r := hitShip.StartPosition.row; r <= hitShip.EndPosition.row; r++ {
			if board[r][hitShip.StartPosition.row] == hit {
				hits++
			}
		}
	}
	// Check if number of hits == ship length
	if hits == shipLength {
		return true, hitShip.ShipName
	}

	return false, ""
}
