package cube_test

import (
	"testing"

	"rubiks-solver/internal/cube"
)

func TestNewSolvedCube(t *testing.T) {
	c := cube.NewSolvedCube()
	for i := 0; i < 9; i++ {
		if c.Facelets[i] != cube.U {
			t.Errorf("Expected U color at position %d", i)
		}
	}
}

func TestApplyMove_U(t *testing.T) {
	c := cube.NewSolvedCube()
	c.ApplyMove(cube.MoveU)

	if c.Facelets[18] != cube.R {
		t.Errorf("Expected Front top row to have Right color (R) after U move")
	}
}

func TestApplyMove_4TimesReturnsToSolved(t *testing.T) {
	c := cube.NewSolvedCube()

	// Act: Apply U move 4 times
	c.ApplyMove(cube.MoveU)
	c.ApplyMove(cube.MoveU)
	c.ApplyMove(cube.MoveU)
	c.ApplyMove(cube.MoveU)

	// Assert: It should be exactly like a newly solved cube
	solved := cube.NewSolvedCube()
	if c.Facelets != solved.Facelets {
		t.Errorf("Cube did not return to solved state after 4 identical moves")
	}
}

func TestApplyMove_PrimeCancelsNormal(t *testing.T) {
	c := cube.NewSolvedCube()

	// Act: R followed by R' (Right inverted)
	c.ApplyMove(cube.MoveR)
	c.ApplyMove(cube.MoveR3) // R3 is R'

	// Assert: They should cancel each other out
	solved := cube.NewSolvedCube()
	if c.Facelets != solved.Facelets {
		t.Errorf("R followed by R' did not return cube to solved state")
	}
}
