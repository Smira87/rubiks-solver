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

func TestSolve_UnsolvableDueToDepthLimit(t *testing.T) {
	// Arrange
	cube.InitMoveTables()
	cube.InitPruningTables()

	c := cube.NewSolvedCubieCube()

	// Apply a 5-move scramble
	scramble := []cube.Move{cube.MoveR, cube.MoveU, cube.MoveF, cube.MoveL, cube.MoveD}
	for _, move := range scramble {
		c.ApplyMove(move)
	}

	// Act: Try to solve it, but give it an impossibly small limit (2 moves)
	solution := solver.Solve(c, 2)

	// Assert: It MUST return nil because 5 random moves can't be reversed in 2 moves
	if solution != nil {
		t.Errorf("Expected nil solution for impossible depth, got %v", solution)
	}
}

func TestSolve_EdgeOrientationHeavyScramble(t *testing.T) {
	// Arrange
	cube.InitMoveTables()
	cube.InitPruningTables()

	c := cube.NewSolvedCubieCube()

	// Apply a scramble that heavily flips edges (Flip) compared to corners (Twist).
	// This forces the A* heuristic to use dist2 > dist1.
	// F, B, L, R moves are known to alter edge orientation.
	scramble := []cube.Move{
		cube.MoveF, cube.MoveB, cube.MoveL, cube.MoveR,
		cube.MoveF2, cube.MoveU, cube.MoveD3,
	}
	for _, move := range scramble {
		c.ApplyMove(move)
	}

	// Act
	solution := solver.Solve(c, 20)

	// Assert
	if solution == nil {
		t.Fatalf("Failed to solve Edge-heavy scramble")
	}

	for _, move := range solution {
		c.ApplyMove(move)
	}

	solved := cube.NewSolvedCubieCube()
	if c.EP != solved.EP || c.EO != solved.EO {
		t.Errorf("Cube edges not properly solved!")
	}
}

func TestSolve_Coverage_Dist2Dominates(t *testing.T) {
	// Arrange
	cube.InitMoveTables()
	cube.InitPruningTables()

	c := cube.NewSolvedCubieCube()

	// White-box testing technique:
	// We deliberately inject a state where Corners and Slice are perfectly solved (dist1 = 0),
	// but Edges are flipped (dist2 > 0).
	// This forces the heuristic A* to take the dist2 path: `minDist = dist2`.
	c.SetFlip(1) // 1 means the first edge is flipped (and the 12th automatically flips to match parity)

	// Act
	// We run the solver. The very first node evaluation will hit our target line.
	solution := solver.Solve(c, 20)

	// Assert
	if solution == nil {
		t.Fatalf("Solver failed on artificial Flip state")
	}
}
