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

func TestSolvePhase1_ReturnsNilWhenMaxDepthExceeded(t *testing.T) {
	// 1. Setup tables
	cube.InitMoveTables()
	cube.InitPruningTables()

	// 2. Create a cube and scramble it well (5 moves)
	c := cube.NewSolvedCubieCube()
	c.ApplyMove(cube.MoveU)
	c.ApplyMove(cube.MoveR)
	c.ApplyMove(cube.MoveF)
	c.ApplyMove(cube.MoveD)
	c.ApplyMove(cube.MoveL)

	// 3. Set an unrealistic goal: solve it in 2 moves maximum
	maxDepth := int8(2)
	solution := solver.SolvePhase1(c, maxDepth)

	// 4. Verify that nil is returned
	if solution != nil {
		t.Errorf("Expected nil because the path is longer than maxDepth, but got: %v", solution)
	}
}

func TestSolvePhase1_Coverage_FlipDistanceGreater(t *testing.T) {
	cube.InitMoveTables()
	cube.InitPruningTables()

	c := cube.NewSolvedCubieCube()

	// Artificially mess up only the edges (Flip), leaving the corners (Twist) solved.
	// This ensures that dist2 (edges) will be greater than dist1 (corners),
	// allowing us to cover the line `minDist = dist2`.
	c.SetFlip(500)

	solver.SolvePhase1(c, 5)
}

func TestSolvePhase1_Coverage_LoopExhaustion(t *testing.T) {
	cube.InitMoveTables()
	cube.InitPruningTables()

	c := cube.NewSolvedCubieCube()

	// Create an artificial conflict of interest to cover the `return false` at the end of the loop.
	// 1. Set corners to a state that is solved by 1 move of U.
	c.SetTwist(cube.TwistMove[0][cube.MoveU])
	// 2. Set edges to a state that is solved by 1 move of F.
	c.SetFlip(cube.FlipMove[0][cube.MoveF])

	// The heuristic (minDist) will say: "Oh, max distance = 1, let's enter the loop!".
	// The algorithm will try move U (corners will be solved, but edges won't -> false).
	// It will try move F (edges will be solved, but corners won't -> false).
	// It will try the remaining 16 moves -> all false.
	// The loop will exhaust, and the function will reach the final `return false`!

	solver.SolvePhase1(c, 1)
}
