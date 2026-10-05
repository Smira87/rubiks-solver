package solver

import (
	"testing"

	"rubiks-solver/internal/cube"
)

// TestSearchPhase1_Coverage artificially forces the DFS loop to hit
// the 'continue' statement (redundant face moves) and the final
// 'return false' statement (loop exhaustion).
func TestSearchPhase1_Coverage(t *testing.T) {
	// 1. Initialize tables
	cube.InitMoveTables()
	cube.InitPruningTables()

	// 2. Create a state that is exactly 1 move away from being solved.
	// We apply F' (MoveF3), so the winning move to solve it is F (MoveF).
	c := cube.NewSolvedCubieCube()
	c.ApplyMove(cube.MoveF3)

	twist := c.GetTwist()
	flip := c.GetFlip()
	slice := c.GetUDSlice()

	path := make([]cube.Move, 1)

	// 3. The Trick: Call the private searchPhase1 function directly.
	// We give it depth=1, but we set lastFace=2 (which is the Front face).
	// What happens next:
	// - The loop will check U, R, D, L, B. They will recurse to depth 0 and fail.
	// - When the loop checks F, F2, F' (m=6,7,8), `face == lastFace` will be true,
	//   and it will hit the `continue` statement (Covering the first red block!).
	// - Since the only winning move (F) was skipped, all 18 iterations will fail.
	// - The loop will finish, hitting `return false` (Covering the second red block!).

	result := searchPhase1(twist, flip, slice, 1, 2, path, 0)

	// Assert
	if result != false {
		t.Errorf("Expected searchPhase1 to return false because the winning move was artificially blocked")
	}
}
