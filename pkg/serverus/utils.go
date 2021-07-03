package serverus

// removeMoveOption package private utility function to remove string from []string
func removeMoveOption(possibleMoves *[]string, move string) {
	moves := *possibleMoves
	for i, m := range moves {
		if m == move {
			moves = append(moves[:i], moves[i+1:]...)
		}
	}
	*possibleMoves = moves
}

func removeMutlipleMoveOptions(possibleMoves *[]string, movesToRemove []string) {
	moveList := *possibleMoves
	var result []string
	for _, move := range moveList {
		if !Contains(movesToRemove, move) {
			result = append(result, move)
		}
	}
	*possibleMoves = result
}

// contains checks if a string is present in a slice
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
