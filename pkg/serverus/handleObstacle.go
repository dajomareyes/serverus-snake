package serverus

var UP Coord = Coord{0, 1}
var DOWN Coord = Coord{0, -1}
var LEFT Coord = Coord{-1, 0}
var RIGHT Coord = Coord{1, 0}

// returns a list of possible moves
func HandleObstacle(gameState GameRequest, possibleMoves []string) []string {
	// splice snake body to exclude the head
	var candidatesForRemoval []string

	for _, move := range possibleMoves {
		if isObstacleSelf(move, gameState.You.Head, gameState.You.Body) {
			candidatesForRemoval = append(candidatesForRemoval, move)
		}
	}

	removeMutlipleMoveOptions(&possibleMoves, candidatesForRemoval)

	return possibleMoves
}

func isObstacleSelf(move string, head Coord, snakeBody []Coord) bool {
	var direction Coord

	switch move {
	case "up":
		direction = UP
	case "down":
		direction = DOWN
	case "left":
		direction = LEFT
	case "right":
		direction = RIGHT
	}

	// create next move
	nextMove := Coord{head.X + direction.X, head.Y + direction.Y}

	// check if next move is in body
	for _, node := range snakeBody {
		if nextMove == node {
			return true
		}
	}

	return false
}
