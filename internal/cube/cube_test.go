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

func TestMoveU(t *testing.T) {
	// Arrange: Create a new solved cube
	c := cube.NewSolvedCube()

	// Act: Apply the Up move (90 degrees clockwise)
	c.MoveU()

	// Assert: Check adjacent faces
	// After U move, the top row of the Front face (indices 18, 19, 20)
	// should contain the colors that were previously on the Right face (cube.R)
	if c.Facelets[18] != cube.R || c.Facelets[19] != cube.R || c.Facelets[20] != cube.R {
		t.Errorf("Expected Front top row to have Right color (R), got %v, %v, %v",
			c.Facelets[18], c.Facelets[19], c.Facelets[20])
	}

	// The top row of the Left face (indices 36, 37, 38)
	// should contain the colors of the Front face (cube.F)
	if c.Facelets[36] != cube.F || c.Facelets[37] != cube.F || c.Facelets[38] != cube.F {
		t.Errorf("Expected Left top row to have Front color (F), got %v, %v, %v",
			c.Facelets[36], c.Facelets[37], c.Facelets[38])
	}
}
