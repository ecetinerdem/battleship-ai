package main

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
			p.heatMap[i][j] = 1

			// Increase Probability in a checker board pattern
			if (i+j)%2 == 0 {
				p.heatMap[i][j] += 1
			}

			// Hihgher probability in the center
			center := abs(i-boardSize/2) + abs(j-boardSize/2)

			if center <= 3 {
				p.heatMap[i][j] += 2
			}
		}
	}
}

func (p *AIPlayer) GetBoard() *Board {
	return &p.board
}
