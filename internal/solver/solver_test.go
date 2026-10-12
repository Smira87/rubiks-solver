package solver_test

import (
	"testing"

	"rubiks-solver/internal/cube"
	"rubiks-solver/internal/solver"
)

func TestSolveFullCube(t *testing.T) {
	// Arrange: Initialize all tables (Pre-computation)
	cube.InitMoveTables()
	cube.InitPruningTables()

	// Create a solved cube
	c := cube.NewSolvedCubieCube()

	// Apply a known chaotic scramble
	// Scramble: R, U', F2, D, L', B2
	scramble := []cube.Move{
		cube.MoveR, cube.MoveU3, cube.MoveF2,
		cube.MoveD, cube.MoveL3, cube.MoveB2,
	}

	for _, move := range scramble {
		c.ApplyMove(move)
	}

	// Act: Solve the cube! Standard Kociemba can solve any cube in <= 22 moves.
	solution := solver.Solve(c, 22)

	// Assert: Check if a solution was found
	if solution == nil {
		t.Fatalf("Solver failed to find a solution within 22 moves")
	}

	// Apply the solution to our scrambled cube
	for _, move := range solution {
		c.ApplyMove(move)
	}

	// Verify that the cube is perfectly solved
	solved := cube.NewSolvedCubieCube()
	if c.CP != solved.CP || c.EP != solved.EP || c.CO != solved.CO || c.EO != solved.EO {
		t.Errorf("Cube is not solved after applying the solution sequence!")
	}

	t.Logf("Successfully solved! Scramble was %d moves, Solution is %d moves.", len(scramble), len(solution))
}
