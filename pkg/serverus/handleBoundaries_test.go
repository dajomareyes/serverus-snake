package serverus

import "testing"

func TestIsAtUpperBoundary(t *testing.T) {
	result := isAtUpperBoundary()

	if result != "hello" {
		t.Errorf("isAtUpperBoundary failed, expected 'hello', got %v", result)
	} else {
		t.Log("isAtUpperBoundary passed")
	}
}
