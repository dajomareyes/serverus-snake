package serverus_test

import (
	"testing"

	. "github.com/dajomareyes/serverus-snake/pkg/serverus"
	"github.com/stretchr/testify/assert"
)

var BOARD_DIMS = Board{Width: 10, Height: 10}

func setup(snakeHead Coord) (GameRequest, []string) {
	return GameRequest{Board: BOARD_DIMS, You: Battlesnake{Head: snakeHead}}, []string{"up", "down", "left", "right"}
}

func TestWhenSnakeIsAtNorthernBoundary(t *testing.T) {
	request, moves := setup(Coord{X: 5, Y: 10})
	HandleBoundaries(request, &moves)
	expected := []string{"down", "left", "right"}
	assert.Equal(t, expected, moves)
}

func TestWhenSnakeIsAtSourthernBoundary(t *testing.T) {
	request, moves := setup(Coord{X: 5, Y: 0})
	HandleBoundaries(request, &moves)
	expected := []string{"up", "left", "right"}
	assert.Equal(t, expected, moves)
}

func TestWhenSnakeIsAtEasternBoundary(t *testing.T) {
	request, moves := setup(Coord{X: 10, Y: 9})
	HandleBoundaries(request, &moves)
	expected := []string{"up", "down", "left"}
	assert.Equal(t, expected, moves)
}

func TestWhenSnakeIsAtWesternBoundary(t *testing.T) {
	request, moves := setup(Coord{X: 0, Y: 5})
	HandleBoundaries(request, &moves)
	expected := []string{"up", "down", "right"}
	assert.Equal(t, expected, moves)
}

func TestHandleBoundariesWhenSnakeIsAtNorthEastCorner(t *testing.T) {
	request, moves := setup(Coord{X: 10, Y: 10})
	HandleBoundaries(request, &moves)
	expected := []string{"down", "left"}
	assert.Equal(t, expected, moves)
}

func TestHandleBoundariesWhenSnakeIsAtNorthWestCorner(t *testing.T) {
	request, moves := setup(Coord{X: 0, Y: 10})
	HandleBoundaries(request, &moves)
	expected := []string{"down", "right"}
	assert.Equal(t, expected, moves)
}

func TestHandleBoundariesWhenSnakeIsAtSouthWestCorner(t *testing.T) {
	request, moves := setup(Coord{X: 0, Y: 10})
	HandleBoundaries(request, &moves)
	expected := []string{"up", "right"}
	assert.Equal(t, expected, moves)
}

func TestHandleBoundariesWhenSnakeIsAtSouthEastCorner(t *testing.T) {
	request, moves := setup(Coord{X: 10, Y: 10})
	HandleBoundaries(request, &moves)
	expected := []string{"up", "left"}
	assert.Equal(t, expected, moves)
}
