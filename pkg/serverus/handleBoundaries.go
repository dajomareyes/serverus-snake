package serverus

// HandleBoundaries removes illegal moves when the next move interacts with a boundary
// Uses the current head location x,y and the boards limits to determine if the next move is possible
// It does not return anything
func HandleBoundaries(gameState GameRequest, moves []string) []string {

	// board dimensions should be offset by -1 to because of 0 counting
	boardWidth := gameState.Board.Width - 1
	boardHeight := gameState.Board.Height - 1
	head := gameState.You.Head
	possibleMoves := moves

	if isAtNorthernBoundary(head, boardHeight) {
		removeMoveOption(&possibleMoves, "up")
	}

	if isAtEasternBoundary(head, boardWidth) {
		removeMoveOption(&possibleMoves, "right")
	}

	if isAtSouthernBoundary(head) {
		removeMoveOption(&possibleMoves, "down")
	}

	if isAtWesternBoundary(head) {
		removeMoveOption(&possibleMoves, "left")
	}

	return possibleMoves
}

func isAtNorthernBoundary(head Coord, boardHeight int) bool {
	nextMove := head.Y + 1
	return nextMove > boardHeight
}

func isAtEasternBoundary(head Coord, boardWidth int) bool {
	nextMove := head.X + 1
	return nextMove > boardWidth
}

func isAtSouthernBoundary(head Coord) bool {
	nextMove := head.Y - 1
	return nextMove < 0
}

func isAtWesternBoundary(head Coord) bool {
	nextMove := head.X - 1
	return nextMove < 0
}
