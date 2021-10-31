package serverus_test

import (
	"fmt"
	"testing"

	. "github.com/dajomareyes/serverus-snake/pkg/serverus"
	"github.com/stretchr/testify/assert"
)

func setupObstacleTest(snakeHead Coord, body []Coord) (GameRequest, []string) {
	var gameState GameRequest = GameRequest {
		You: Battlesnake {
			Head: snakeHead,
			Body: body,
		},
		Board: Board {
			Snakes: []Battlesnake{
				Battlesnake {
					Body: body,
				},
			},
		},
	}
	return gameState, []string{"up", "down", "left", "right"}
}

func TestHandleObstacle(t *testing.T) {
	var obstacleTests = []struct {
		title    string
		head     Coord
		body     []Coord
		expected []string
	}{
		{
			title:    "test when snake body is over the head",
			head:     Coord{X: 5, Y: 5},
			body:     []Coord{{X: 5, Y: 5}, {X: 6, Y: 5}, {X: 6, Y: 6}, {X: 5, Y: 6}, {X: 4, Y: 6}},
			expected: []string{"down", "left"},
		},
		{
			title:    "test when snake body is under the head",
			head:     Coord{X: 5, Y: 5},
			body:     []Coord{{X: 5, Y: 5}, {X: 6, Y: 5}, {X: 6, Y: 4}, {X: 5, Y: 4}, {X: 4, Y: 4}},
			expected: []string{"up", "left"},
		},
		{
			title:    "test when snake body is left the head",
			head:     Coord{X: 5, Y: 5},
			body:     []Coord{{X: 5, Y: 5}, {X: 5, Y: 6}, {X: 4, Y: 6}, {X: 4, Y: 5}, {X: 4, Y: 4}},
			expected: []string{"down", "right"},
		},
		{
			title:    "test when snake body is right of the head",
			head:     Coord{X: 5, Y: 5},
			body:     []Coord{{X: 5, Y: 5}, {X: 5, Y: 6}, {X: 6, Y: 6}, {X: 6, Y: 5}, {X: 6, Y: 4}},
			expected: []string{"down", "left"},
		},
	}

	for _, tt := range obstacleTests {
		t.Run(fmt.Sprint(tt.title), func(t *testing.T) {
			request, moves := setupObstacleTest(tt.head, tt.body)
			result := HandleObstacle(request, moves)
			assert.Equal(t, tt.expected, result)
		})
	}

}
