package cube

// TwistMove caches the transition for the corner orientation coordinate.
// First index is the Twist value (0-2186), second is the Move (0-17).
var TwistMove [2187][18]uint16

// FlipMove caches the transition for the edge orientation coordinate.
// First index is the Flip value (0-2047), second is the Move (0-17).
var FlipMove [2048][18]uint16

// UDSliceMove caches the transition for the UD-slice coordinate.
// First index is the UDSlice value (0-494), second is the Move (0-17).
var UDSliceMove [495][18]uint16

// InitMoveTables pre-calculates the transition tables for Phase 1.
// It must be called once when the application starts.

// InitMoveTables pre-calculates the transition tables for Phase 1.
// It must be called once when the application starts.
func InitMoveTables() {
	// 1. Generate Twist Move Table
	for i := uint16(0); i < 2187; i++ {
		for m := Move(0); m < 18; m++ {
			c := NewSolvedCubieCube()
			c.SetTwist(i)
			c.ApplyMove(m)
			TwistMove[i][m] = c.GetTwist()
		}
	}

	// 2. Generate Flip Move Table
	for i := uint16(0); i < 2048; i++ {
		for m := Move(0); m < 18; m++ {
			c := NewSolvedCubieCube()
			c.SetFlip(i)
			c.ApplyMove(m)
			FlipMove[i][m] = c.GetFlip()
		}
	}

	// 3. Generate UDSlice Move Table
	for i := uint16(0); i < 495; i++ {
		for m := Move(0); m < 18; m++ {
			c := NewSolvedCubieCube()
			c.SetUDSlice(i)
			c.ApplyMove(m)
			UDSliceMove[i][m] = c.GetUDSlice()
		}
	}
}

// SliceTwistPrun is the pruning table for the combination of UDSlice and Twist coordinates.
// We use a 1D array for cache locality and int8 to save memory.
var SliceTwistPrun [495 * 2187]int8

// SliceFlipPrun is the pruning table for the combination of UDSlice and Flip coordinates.
// It provides the minimum number of moves to solve these coordinates in Phase 1.
var SliceFlipPrun [495 * 2048]int8

// InitPruningTables generates the heuristic tables using Breadth-First Search (BFS).
func InitPruningTables() {
	// 1. Initialize all distances to -1 (meaning "unvisited")
	for i := 0; i < len(SliceTwistPrun); i++ {
		SliceTwistPrun[i] = -1
	}
	for i := 0; i < len(SliceFlipPrun); i++ {
		SliceFlipPrun[i] = -1
	}

	// 2. Generate SliceTwistPrun table
	SliceTwistPrun[0] = 0 // The solved state is at distance 0
	depth := int8(0)
	done := 1

	for done < len(SliceTwistPrun) {
		index := 0 // We track the 1D index manually without division
		for slice := uint16(0); slice < 495; slice++ {
			for twist := uint16(0); twist < 2187; twist++ {
				if SliceTwistPrun[index] == depth {
					for m := Move(0); m < 18; m++ {
						newSlice := UDSliceMove[slice][m]
						newTwist := TwistMove[twist][m]
						newIndex := int(newSlice)*2187 + int(newTwist)

						if SliceTwistPrun[newIndex] == -1 {
							SliceTwistPrun[newIndex] = depth + 1
							done++
						}
					}
				}
				index++ // Cheap addition instead of expensive division!
			}
		}
		depth++
	}

	// 3. Generate SliceFlipPrun table
	SliceFlipPrun[0] = 0
	depth = int8(0)
	done = 1

	for done < len(SliceFlipPrun) {
		index := 0
		for slice := uint16(0); slice < 495; slice++ {
			for flip := uint16(0); flip < 2048; flip++ {
				if SliceFlipPrun[index] == depth {
					for m := Move(0); m < 18; m++ {
						newSlice := UDSliceMove[slice][m]
						newFlip := FlipMove[flip][m]
						newIndex := int(newSlice)*2048 + int(newFlip)

						if SliceFlipPrun[newIndex] == -1 {
							SliceFlipPrun[newIndex] = depth + 1
							done++
						}
					}
				}
				index++
			}
		}
		depth++
	}
}
