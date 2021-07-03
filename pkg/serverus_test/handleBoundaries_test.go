package serverus_test

import (
	"testing"

	. "github.com/dajomareyes/serverus-snake/pkg/serverus"
)

var BOARD_DIMS = Board{Width: 10, Height: 10}

func setup(snakeHead Coord) (GameRequest, []string) {
	return GameRequest{Board: BOARD_DIMS, You: Battlesnake{Head: snakeHead}}, []string{"up", "down", "left", "right"}
}

func TestWhenSnakeIsAtNorthernBoundary(t *testing.T) {
	request, moves := setup(Coord{X: 5, Y: 10})
	HandleBoundaries(request, &moves)

	if Contains(moves, "up") {
		t.Errorf("TestWhenSnakeIsAtNorthernBoundary failed, expected not to find 'up', got %v", moves)
	} else {
		t.Log("TestWhenSnakeIsAtNorthernBoundary passed")
	}
}

func TestWhenSnakeIsAtSourthernBoundary(t *testing.T) {
	request, moves := setup(Coord{X: 5, Y: 0})
	HandleBoundaries(request, &moves)

	if Contains(moves, "down") {
		t.Errorf("TestWhenSnakeIsAtNorthernBoundary failed, expected not to find 'down', got %v", moves)
	} else {
		t.Log("TestWhenSnakeIsAtNorthernBoundary passed")
	}
}

func TestWhenSnakeIsAtEasternBoundary(t *testing.T) {
	request, moves := setup(Coord{X: 10, Y: 9})
	HandleBoundaries(request, &moves)

	if Contains(moves, "right") {
		t.Errorf("TestWhenSnakeIsAtNorthernBoundary failed, expected not to find 'right', got %v", moves)
	} else {
		t.Log("TestWhenSnakeIsAtNorthernBoundary passed")
	}
}

func TestWhenSnakeIsAtWesternBoundary(t *testing.T) {
	request, moves := setup(Coord{X: 0, Y: 5})
	HandleBoundaries(request, &moves)

	if Contains(moves, "left") {
		t.Errorf("TestWhenSnakeIsAtNorthernBoundary failed, expected not to find 'left', got %v", moves)
	} else {
		t.Log("TestWhenSnakeIsAtNorthernBoundary passed")
	}
}

func TestHandleBoundariesWhenSnakeIsAtNorthEastCorner(t *testing.T) {
	request, moves := setup(Coord{X: 10, Y: 10})
	HandleBoundaries(request, &moves)

	if Contains(moves, "up") || Contains(moves, "right") {
		t.Errorf("TestWhenSnakeIsAtNorthernBoundary failed, expected not to find 'up' or 'right', got %v", moves)
	} else {
		t.Log("TestWhenSnakeIsAtNorthernBoundary passed")
	}
}
