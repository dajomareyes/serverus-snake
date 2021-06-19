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
