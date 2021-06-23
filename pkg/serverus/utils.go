package serverus

import "fmt"

// removeMoveOption package private utility function to remove string from []string
func removeMoveOption(possibleMoves *[]string, move string) {
	fmt.Printf("move list - %s removing move option: %s", possibleMoves, move)
	moves := *possibleMoves
	for i, m := range moves {
		if m == move {
			moves = append(moves[:i], moves[i+1:]...)
		}
	}
	*possibleMoves = moves
	fmt.Printf("--> new move list %s\n", possibleMoves)
}

func removeMutlipleMoveOptions(possibleMoves *[]string, movesToRemove []string) {
	fmt.Printf("moves to remove candidates %s ?? ", movesToRemove)
	moveList := *possibleMoves
	var result []string
	for _, move := range moveList {
		if !contains(movesToRemove, move) {
			result = append(result, move)
		}
	}
	fmt.Printf("old move list %s --> new move list %s\n", moveList, result)
	*possibleMoves = result
}

// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
