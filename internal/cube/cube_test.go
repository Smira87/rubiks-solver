package cube_test

import (
	"testing"

	"rubiks-solver/internal/cube"
)

func TestNewSolvedCube(t *testing.T) {
	// Arrange
	c := cube.NewSolvedCube()

	// Act & Assert
	// Check if the first 9 facelets belong to the U (Up) face
	for i := 0; i < 9; i++ {
		if c.Facelets[i] != cube.U {
			t.Errorf("Expected color U at position %d, got %v", i, c.Facelets[i])
		}
	}

	// Check the first facelet of the R (Right) face (position 9)
	if c.Facelets[9] != cube.R {
		t.Errorf("Expected color R at position 9, got %v", c.Facelets[9])
	}
}
