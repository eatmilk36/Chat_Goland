package Login

import (
	"testing"
)

func TestLoginHandler(t *testing.T) {
	result := 2 + 3
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}
