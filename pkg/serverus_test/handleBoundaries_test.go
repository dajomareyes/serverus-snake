package serverus_test

import (
	"fmt"
	"testing"

	. "github.com/dajomareyes/serverus-snake/pkg/serverus"
	"github.com/stretchr/testify/assert"
)

var BOARD_DIMS = Board{Width: 10, Height: 10}

func setup(snakeHead Coord) (GameRequest, []string) {
	return GameRequest{Board: BOARD_DIMS, You: Battlesnake{Head: snakeHead}}, []string{"up", "down", "left", "right"}
}

func TestHandleBoundaries(t *testing.T) {
	var boundaryTests = []struct {
		title    string
		head     Coord
		expected []string
	}{
		{
			title:    "test when snake is at the northen boundary",
			head:     Coord{X: 4, Y: 9},
			expected: []string{"down", "left", "right"},
		},
		{
			title:    "test when snake is at the southern boundary",
			head:     Coord{X: 4, Y: 0},
			expected: []string{"up", "left", "right"},
		},
		{
			title:    "test when snake is at the eastern boundary",
			head:     Coord{X: 9, Y: 4},
			expected: []string{"up", "down", "left"},
		},
		{
			title:    "test when snake is at the western boundary",
			head:     Coord{X: 0, Y: 4},
			expected: []string{"up", "down", "right"},
		},
		{
			title:    "test when snake is at the northeast corner",
			head:     Coord{X: 9, Y: 9},
			expected: []string{"down", "left"},
		},
		{
			title:    "test when snake is at the northwest corner",
			head:     Coord{X: 0, Y: 9},
			expected: []string{"down", "right"},
		},
		{
			title:    "test when snake is at the southeast corner",
			head:     Coord{X: 9, Y: 0},
			expected: []string{"up", "left"},
		},
		{
			title:    "test when snake is at the southwest corner",
			head:     Coord{X: 0, Y: 0},
			expected: []string{"up", "right"},
		},
	}

	for _, tt := range boundaryTests {
		t.Run(fmt.Sprint(tt.title), func(t *testing.T) {
			t.Parallel()
			result := HandleBoundaries(setup(tt.head))
			assert.Equal(t, tt.expected, result)
		})
	}
}
