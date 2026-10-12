package solver

import (
	"rubiks-solver/internal/cube"
)

// Phase2Moves contains the 10 allowable moves in Phase 2.
var phase2Moves = []cube.Move{
	cube.MoveU, cube.MoveU2, cube.MoveU3,
	cube.MoveD, cube.MoveD2, cube.MoveD3,
	cube.MoveR2, cube.MoveL2, cube.MoveF2, cube.MoveB2,
}

// Solve is the main entry point for the Kociemba Two-Phase Algorithm.
// It returns a slice of moves to perfectly solve the given CubieCube.
func Solve(c *cube.CubieCube, maxTotalMoves int8) []cube.Move {
	twist := c.GetTwist()
	flip := c.GetFlip()
	slice := c.GetUDSlice()

	// Allocate a slice to hold our potential solution path.
	// We allocate the maximum possible size once to avoid allocations during DFS.
	path := make([]cube.Move, maxTotalMoves)

	// Phase 1 Iterative Deepening
	for p1Depth := int8(0); p1Depth <= maxTotalMoves; p1Depth++ {
		// We pass the original cube down so we can calculate Phase 2 coords later
		solutionLen := searchPhase1(c, twist, flip, slice, p1Depth, maxTotalMoves, -1, path, 0)
		if solutionLen > 0 {
			// A complete Phase 1 + Phase 2 solution was found!
			return path[:solutionLen]
		}
	}

	return nil // Unsolvable within maxTotalMoves
}

// searchPhase1 explores Phase 1. When it reaches a Phase 1 solution, it triggers searchPhase2.
func searchPhase1(
	originalCube *cube.CubieCube,
	twist, flip, slice uint16,
	p1Depth, maxTotalMoves int8,
	lastFace int,
	path []cube.Move,
	pathLen int,
) int {

	// 1. We reached the end of our Phase 1 depth limit
	if p1Depth == 0 {
		if twist == 0 && flip == 0 && slice == 0 {
			// --- TRANSITION TO PHASE 2 ---

			// We found a Phase 1 solution. Now, what is the state of the cube?
			// Let's create a copy and apply the Phase 1 path to it.
			tempCube := cube.GetPooledCubie()

			// Copy the original state
			tempCube.CP = originalCube.CP
			tempCube.CO = originalCube.CO
			tempCube.EP = originalCube.EP
			tempCube.EO = originalCube.EO

			// Apply Phase 1 moves
			for i := 0; i < pathLen; i++ {
				tempCube.ApplyMove(path[i])
			}

			// Extract Phase 2 coordinates
			cPerm := tempCube.GetCPerm()
			ePerm := tempCube.GetEPerm()
			mPerm := tempCube.GetMPerm()

			cube.ReleaseCubie(tempCube) // Return to pool

			// Phase 2 Iterative Deepening
			// How many moves do we have left in our "budget"?
			movesLeft := maxTotalMoves - int8(pathLen)

			for p2Depth := int8(0); p2Depth <= movesLeft; p2Depth++ {
				finalLen := searchPhase2(cPerm, ePerm, mPerm, p2Depth, lastFace, path, pathLen)
				if finalLen > 0 {
					return finalLen // Solved!
				}
			}
		}
		return 0 // This branch failed
	}

	// 2. A* Pruning for Phase 1
	dist1 := cube.SliceTwistPrun[int(slice)*2187+int(twist)]
	dist2 := cube.SliceFlipPrun[int(slice)*2048+int(flip)]
	minDist := dist1
	if dist2 > minDist {
		minDist = dist2
	}
	if minDist > p1Depth {
		return 0 // Prune branch
	}

	// 3. Branching (18 moves)
	for m := cube.Move(0); m < 18; m++ {
		face := int(m) / 3
		if face == lastFace {
			continue // Skip redundant faces (e.g., U followed by U2)
		}

		path[pathLen] = m

		newTwist := cube.TwistMove[twist][m]
		newFlip := cube.FlipMove[flip][m]
		newSlice := cube.UDSliceMove[slice][m]

		res := searchPhase1(originalCube, newTwist, newFlip, newSlice, p1Depth-1, maxTotalMoves, face, path, pathLen+1)
		if res > 0 {
			return res
		}
	}

	return 0
}

// searchPhase2 is the DFS guided by Phase 2 pruning tables.
func searchPhase2(
	cPerm, ePerm, mPerm uint16,
	depth int8,
	lastFace int,
	path []cube.Move,
	pathLen int,
) int {

	// 1. Base case: Check if solved
	if depth == 0 {
		if cPerm == 0 && ePerm == 0 && mPerm == 0 {
			return pathLen // Return the total length of the solution
		}
		return 0
	}

	// 2. A* Pruning for Phase 2
	dist1 := cube.MCPermPrun[int(mPerm)*40320+int(cPerm)]
	dist2 := cube.MEPermPrun[int(mPerm)*40320+int(ePerm)]
	minDist := dist1
	if dist2 > minDist {
		minDist = dist2
	}
	if minDist > depth {
		return 0 // Prune
	}

	// 3. Branching (ONLY 10 Phase 2 moves!)
	for _, m := range phase2Moves {
		face := int(m) / 3
		if face == lastFace {
			continue
		}

		path[pathLen] = m

		newCPerm := cube.CPermMove[cPerm][m]
		newEPerm := cube.EPermMove[ePerm][m]
		newMPerm := cube.MPermMove[mPerm][m]

		res := searchPhase2(newCPerm, newEPerm, newMPerm, depth-1, face, path, pathLen+1)
		if res > 0 {
			return res
		}
	}

	return 0
}
