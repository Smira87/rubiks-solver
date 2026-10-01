package cube_test

import (
	"testing"

	"rubiks-solver/internal/cube"
)

func TestCnk(t *testing.T) {
	// Arrange: Table-driven tests (Best Practice in Go)
	tests := []struct {
		n        int
		k        int
		expected int
	}{
		{n: 12, k: 4, expected: 495},
		{n: 4, k: 0, expected: 1},
		{n: 4, k: 4, expected: 1},
		{n: 5, k: 2, expected: 10},
	}

	// Act & Assert
	for _, tt := range tests {
		actual := cube.Cnk(tt.n, tt.k)
		if actual != tt.expected {
			t.Errorf("Cnk(%d, %d) = %d; expected %d", tt.n, tt.k, actual, tt.expected)
		}
	}
}
