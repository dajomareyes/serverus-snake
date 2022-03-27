package serverus

var UP Coord = Coord{0, 1}
var DOWN Coord = Coord{0, -1}
var LEFT Coord = Coord{-1, 0}
var RIGHT Coord = Coord{1, 0}

// HandleObstacle returns a list of possible moves
func HandleObstacle(gameState GameRequest, possibleMoves []string) []string {
	// splice snake body to exclude the head
	var candidatesForRemoval []string

	for _, move := range possibleMoves {
		if isObstacleSnake(move, gameState.You.Head, gameState.Board.Snakes) {
			candidatesForRemoval = append(candidatesForRemoval, move)
		}
	}

	removeMutlipleMoveOptions(&possibleMoves, candidatesForRemoval)

	return possibleMoves
}

func isObstacleSnake(move string, head Coord, snakes []Battlesnake) bool {
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

  nextMove := Coord{head.X + direction.X, head.Y + direction.Y}

  // for every battle snake in the game make sure the next move doesn't run
  // into a part of their snake
  for _, snake := range snakes {
    for _, snakeBodyCoordinate := range snake.Body {
      if nextMove == snakeBodyCoordinate {
        // exit out of this for loop this move won't work
        return true
      }
    }
  }

  return false
}
