package serverus

import "fmt"

var UP Coord = Coord{0, 1}
var DOWN Coord = Coord{0, -1}
var LEFT Coord = Coord{-1, 0}
var RIGHT Coord = Coord{1, 0}

func HandleObstacle(gameState GameRequest, possibleMoves *[]string) {
	// splice snake body to exclude the head
	fmt.Println("IN HandleObstacle")
	snakeBody := gameState.You.Body

	var candidatesForRemoval []string

	for _, move := range *possibleMoves {
		if isObstacleSelf(move, gameState.You.Head, snakeBody) {
			candidatesForRemoval = append(candidatesForRemoval, move)
		}
	}

	removeMutlipleMoveOptions(possibleMoves, candidatesForRemoval)
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
	fmt.Printf("next move %s direction (%d): %d, snake body [%d]\n", move, direction, nextMove, snakeBody)
	for _, node := range snakeBody {
		if nextMove == node {
			return true
		}
	}

	return false
}
