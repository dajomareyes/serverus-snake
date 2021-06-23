package serverus

import "fmt"

var boardWidth int
var boardHeight int

// HandleBoundaries removes illegal moves when the next move interacts with a boundary
// Uses the current head location x,y and the boards limits to determine if the next move is possible
// It does not return anything
func HandleBoundaries(gameState GameRequest, possibleMoves *[]string) {
	fmt.Println("IN HandleBoundaries")

	// board dimensions should be offset by -1 to take into account 0 counting
	boardWidth = gameState.Board.Width - 1
	boardHeight = gameState.Board.Height - 1

	fmt.Println(boardHeight, boardWidth)

	head := gameState.You.Head

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

func isAtNorthernBoundary(head Coord) bool {
	nextMove := head.Y + 1
	return nextMove > boardHeight
}

func isAtSouthernBoundary(head Coord) bool {
	nextMove := head.Y - 1
	return nextMove < 0
}

func isAtWesternBoundary(head Coord) bool {
	nextMove := head.X - 1
	return nextMove < 0
}

func isAtEasternBoundary(head Coord) bool {
	nextMove := head.X + 1
	return nextMove > boardWidth
}
