package solver

import (
	"rubiks-solver/internal/cube"
)

// SolvePhase1 finds a sequence of moves to bring the cube into the G1 group
// (where Twist, Flip, and UDSlice are all 0).
func SolvePhase1(c *cube.CubieCube, maxDepth int8) []cube.Move {
	// We extract the initial coordinates.
	// The search engine only works with integers (coordinates), not the full cube!
	twist := c.GetTwist()
	flip := c.GetFlip()
	slice := c.GetUDSlice()

	// Iterative Deepening: try to find a solution in 0 moves, then 1, then 2, etc.
	for depth := int8(0); depth <= maxDepth; depth++ {
		// path will store the current sequence of moves being explored
		path := make([]cube.Move, depth)

		// -1 for lastFace means "no previous move"
		if searchPhase1(twist, flip, slice, depth, -1, path, 0) {
			return path
		}
	}

	return nil // No solution found within maxDepth
}

// searchPhase1 is the recursive DFS function guided by the heuristic pruning tables (A*).
func searchPhase1(twist, flip, slice uint16, depth int8, lastFace int, path []cube.Move, pathLen int) bool {
	// 1. Base case: we reached the target depth. Check if solved.
	if depth == 0 {
		return twist == 0 && flip == 0 && slice == 0
	}

	// 2. A* Pruning (The Magic)
	// We ask our pre-computed tables: "What is the absolute minimum number of moves
	// required to solve the Twist+Slice and Flip+Slice from here?"
	dist1 := cube.SliceTwistPrun[int(slice)*2187+int(twist)]
	dist2 := cube.SliceFlipPrun[int(slice)*2048+int(flip)]

	// The true distance is at least the maximum of the two estimates.
	minDist := dist1
	if dist2 > minDist {
		minDist = dist2
	}

	// If the minimum required moves exceed the moves we have left (depth),
	// this branch is a dead end. Prune it!
	if minDist > depth {
		return false
	}

	// 3. Explore all 18 moves
	for m := cube.Move(0); m < 18; m++ {
		// Smart Branching: Avoid redundant moves
		// e.g., if we just turned the U face (m/3 == 0), don't turn U, U2, or U' next.
		face := int(m) / 3
		if face == lastFace {
			continue
		}

		// Calculate new coordinates using our O(1) Move Tables
		newTwist := cube.TwistMove[twist][m]
		newFlip := cube.FlipMove[flip][m]
		newSlice := cube.UDSliceMove[slice][m]

		// Record the move in our path
		path[pathLen] = m

		// Dive deeper (depth decreases by 1)
		if searchPhase1(newTwist, newFlip, newSlice, depth-1, face, path, pathLen+1) {
			return true
		}
	}

	return false
}
