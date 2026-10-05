package solver_test

import (
	"testing"

	"rubiks-solver/internal/cube"
	"rubiks-solver/internal/solver"
)

func TestSolvePhase1(t *testing.T) {
	// Arrange: Initialize tables (this takes a moment, but only happens once)
	cube.InitMoveTables()
	cube.InitPruningTables()

	// Create a solved cube and apply a known scramble (e.g., U, R, F)
	c := cube.NewSolvedCubieCube()
	c.ApplyMove(cube.MoveU)
	c.ApplyMove(cube.MoveR)
	c.ApplyMove(cube.MoveF)

	// Act: Ask the solver to find a path back to Phase 1 target
	// The maximum depth is set to 12 (Phase 1 rarely takes more than 12 moves)
	solution := solver.SolvePhase1(c, 12)

	// Assert
	if solution == nil {
		t.Fatalf("Solver failed to find a solution for a 3-move scramble")
	}

	// Apply the solution to the scrambled cube
	for _, move := range solution {
		c.ApplyMove(move)
	}

	// Verify that Phase 1 is strictly solved (Coordinates must be 0)
	if twist := c.GetTwist(); twist != 0 {
		t.Errorf("Phase 1 not solved: Twist is %d, expected 0", twist)
	}
	if flip := c.GetFlip(); flip != 0 {
		t.Errorf("Phase 1 not solved: Flip is %d, expected 0", flip)
	}
	if slice := c.GetUDSlice(); slice != 0 {
		t.Errorf("Phase 1 not solved: UDSlice is %d, expected 0", slice)
	}
}
