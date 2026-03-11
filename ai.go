package main

import (
	"fmt"
	"math/rand"
)

const (
	// Heatmap weights
	baseProbability   = 1   // Starting probability for valid cells
	checkerBoardBonus = 1   // Bonus for the checkerboard pattern (less likeliy adjacent placements)
	centerProximity   = 2   // Bonus for cells closer to the center
	maxCenterDistance = 3   // How far from the center qualifies for the bonus
	huntModeBoost     = 100 // Significant bonus for cells adjacent hits in hunt mode
	shipFitBonus      = 2   // Base bonus multiplier for fitting a ship
)

type AIPlayer struct {
	board         Board
	heatMap       [boardSize][boardSize]int
	hits          []Position
	shipsSunk     int
	huntMode      bool
	potentilShips []struct {
		size    int
		sunk    bool
		hits    int
		shipPos []Position
	}
	ships    []Ship
	opponent *HumanPlayer
}

func NewAIPlayer() *AIPlayer {
	p := &AIPlayer{
		shipsSunk: 0,
		huntMode:  false,
	}

	for i := range boardSize {
		for j := range boardSize {
			p.board[i][j] = empty
		}
	}

	//Initialize heat map
	p.initializeHeatMap()

	// Initialize potential ship tracking
	p.potentilShips = make([]struct {
		size    int
		sunk    bool
		hits    int
		shipPos []Position
	}, len(shipTypes))

	for i, shipType := range shipTypes {
		p.potentilShips[i].size = shipType.size
		p.potentilShips[i].sunk = false
		p.potentilShips[i].hits = 0
		p.potentilShips[i].shipPos = make([]Position, 0)
	}
	return p
}

func (p *AIPlayer) initializeHeatMap() {
	for i := range boardSize {
		for j := range boardSize {
			// Start with base probability
			p.heatMap[i][j] = baseProbability

			// Increase Probability in a checker board pattern
			if (i+j)%2 == 0 {
				p.heatMap[i][j] += checkerBoardBonus
			}

			// Hihgher probability in the center
			center := abs(i-boardSize/2) + abs(j-boardSize/2)

			if center <= maxCenterDistance {
				p.heatMap[i][j] += centerProximity
			}
		}
	}
}

// UpdateHeatmap recalculates the probability heat map based on the current game state
// It considers potential ship placements and priorities targets during huntmode
func (p *AIPlayer) updateHeatMap(opponentBoard *Board) {
	// Reset heatmap and clear previous probabilities
	for i := range boardSize {
		for j := range boardSize {
			p.heatMap[i][j] = 0
		}
	}

	// Calculate base probability
	for r := range boardSize {
		for c := range boardSize {
			// Skip cells that have been already targeted
			if opponentBoard[r][c] == hit || opponentBoard[r][c] == miss {
				continue
			}

			// Assign base probability for valid untargeted cells
			p.heatMap[r][c] = baseProbability

			// Iterate through opponents ships that have not been sunk yet
			for _, shipData := range p.potentilShips {
				if shipData.sunk {
					continue
				}

				shipSize := shipData.size

				// Check horizantal fit: can and unsunk ship of this size fit horizontally starting here?
				if c+shipSize <= boardSize {
					canFitHorizantal := true
					for k := range shipSize {
						// Check if any cell needed for the ship is already a miss or a hit
						if opponentBoard[r][c+k] == miss || opponentBoard[r][c+k] == hit {
							canFitHorizantal = false
							break
						}

						if canFitHorizantal {
							// Increase probability based on ship size if it fits
							p.heatMap[r][c] += shipFitBonus * shipSize
						}
					}
				}
				//Check vertical fit

				if r+shipSize <= boardSize {
					canFitVertical := true
					for k := range shipSize {
						// Check if any cell needed for the ship is already a miss or a hit
						if opponentBoard[r+k][c] == miss || opponentBoard[r+k][c] == hit {
							canFitVertical = false
							break
						}

						if canFitVertical {
							// Increase probability based on ship size if it fits
							p.heatMap[r][c] += shipFitBonus * shipSize
						}
					}
				}
			}
		}
	}
	// Apply huntmode boosts if applicable
	if p.huntMode && len(p.hits) > 0 {
		p.applyHuntModeBoosts(opponentBoard)
	}

}

func (p *AIPlayer) applyHuntModeBoosts(opponentBoard *Board) {

}

func (p *AIPlayer) TakeTurn(opponentBoard *Board) (Position, bool) {
	fmt.Println("\nAI is taking its turn...")

	if p.huntMode {
		fmt.Println("AI is hunt mode!")

	} else {
		fmt.Println("AI is probability target mode!")

	}

	// Update heat map based on game state

	// Select a target based on strategy

	// if in hunt mode...
	// // find the highest probability cells
	// // select a random target from the highest probability cells
	// // if can't find one fall back to random targeting

	// ...else do something else
	// // find the highest probability cells
	// // select a random target from the  cells

	// Perform attack
	// Check to see if we hit
	// if we hit enter hunt mode
	// check to see if ship was sunk
}

func (p *AIPlayer) GetBoard() *Board {
	return &p.board
}

func (p *AIPlayer) PlaceShips() {
	// Place ships in a mix of edge of center clusters

	// Attempt to place larger ships near the edges
	for i, shipType := range shipTypes {
		placed := false
		attempts := 0

		for !placed && attempts < 100 {
			attempts++

			// Decide on placement strategy based on shipsize

			var row, col int

			horizontal := rand.Intn(2) == 0
			if shipType.size >= 4 {
				if horizontal {
					row = rand.Intn(boardSize)
					if rand.Intn(2) == 0 {
						// Near the left edge
						col = rand.Intn(3)
					} else {
						// Near right edge
						col = boardSize - shipType.size - rand.Intn(3)
					}
				} else {
					col = rand.Intn(boardSize)
					if rand.Intn(2) == 0 {
						// Near the top edge
						row = rand.Intn(3)
					} else {
						// Near the bottom edge
						row = boardSize - shipType.size - rand.Intn(3)
					}
				}
			} else {
				// Place smaller ships in a more distributed pattern
				if horizontal {
					row = rand.Intn(boardSize)
					col = rand.Intn(boardSize - shipType.size + 1)
				} else {
					row = rand.Intn(boardSize - shipType.size + 1)
					col = rand.Intn(boardSize)
				}
			}

			// Check if placement is valid
			positions := []Position{}
			valid := true

			for j := range shipType.size {
				var r, c int
				if horizontal {
					r, c = row, col+j
				} else {
					r, c = row+j, col
				}

				// Check validity
				if r < 0 || r >= boardSize || c < 0 || c >= boardSize || p.board[r][c] == ship {
					valid = false
					break
				}

				// Check surrounding cells to aviod placing ships adjacent to one another

				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						nr, nc := r+dr, c+dc
						if nr >= 0 && nr < boardSize && nc >= 0 && nc < boardSize && p.board[nr][nc] == ship && (dr != 0 || dc != 0) {
							// Avoid placing ships diagonally or directly adjacent to another ship
							if i < 2 { // Only for larger ships
								valid = false
								break
							}
						}
					}
					if !valid {
						break
					}
				}
				if !valid {
					break
				}
				positions = append(positions, Position{r, c})
			}
			if valid {
				// Place the ship
				for _, pos := range positions {
					p.board[pos.row][pos.col] = ship
				}

				// Append this ship to the slice of ai ships
				newShip := Ship{
					StartPosition: positions[0],
					EndPosition:   positions[len(positions)-1],
					ShipName:      shipType.name,
				}

				p.ships = append(p.ships, newShip)
				placed = true
			}

		}
		// If we couldn't place the ship, fall back to random placement

		if !placed {
			for {
				horizontal := rand.Intn(2) == 0
				var row, col int

				if horizontal {
					row = rand.Intn(boardSize)
					col = rand.Intn(boardSize - shipType.size + 1)
				} else {
					row = rand.Intn(boardSize - shipType.size + 1)
					col = rand.Intn(boardSize)
				}

				// Check if placement is valid
				valid := true
				positions := []Position{}

				for i := range shipType.size {
					var r, c int
					if horizontal {
						r, c = row, col+i
					} else {
						r, c = row+i, col
					}

					if p.board[r][c] == ship {
						valid = false
						break
					}

					positions = append(positions, Position{r, c})
				}

				if valid {
					// Place ship
					for _, pos := range positions {
						p.board[pos.row][pos.col] = ship
					}

					// Append to slice of ai ship
					newShip := Ship{
						StartPosition: positions[0],
						EndPosition:   positions[len(positions)-1],
					}
					p.ships = append(p.ships, newShip)
					break
				}

			}
		}
	}
}
