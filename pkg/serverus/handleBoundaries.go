package serverus

var boardWidth int
var boardHeight int

// HandleBoundaries removes illegal moves when the next move interacts with a boundary
// Uses the current head location x,y and the boards limits to determine if the next move is possible
// It does not return anything
func HandleBoundaries(gameState GameRequest, possibleMoves *[]string) {
	// board dimensions should be offset by -1 to take into account 0 counting
	boardWidth = gameState.Board.Width - 1
	boardHeight = gameState.Board.Height - 1

	head := []int{gameState.You.Head.X, gameState.You.Head.Y}

	if isAtNorthernBoundary(head) {
		removeMoveOption(possibleMoves, "up")
	}

	if isAtSouthernBoundary(head) {
		removeMoveOption(possibleMoves, "down")
	}

	if isAtWesternBoundary(head) {
		removeMoveOption(possibleMoves, "left")
	}

	if isAtEasternBoundary(head) {
		removeMoveOption(possibleMoves, "right")
	}
}

func isAtNorthernBoundary(head []int) bool {
	nextMove := head[1] + 1
	if nextMove == boardHeight {
		return true
	}
	return false
}

func isAtSouthernBoundary(head []int) bool {
	nextMove := head[1] - 1
	if nextMove == 0 {
		return true
	}
	return false
}

func isAtWesternBoundary(head []int) bool {
	nextMove := head[0] - 1
	if nextMove == 0 {
		return true
	}
	return false
}

func isAtEasternBoundary(head []int) bool {
	nextMove := head[0] + 1
	if nextMove == boardWidth {
		return true
	}
	return false
}

func removeMoveOption(possibleMoves *[]string, move string) {
	moves := *possibleMoves
	for i, m := range moves {
		if m == move {
			moves = append(moves[:i], moves[i+1:]...)
		}
	}
	*possibleMoves = moves
}
