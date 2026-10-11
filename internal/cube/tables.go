package cube

// Phase 1 Move Tables

// TwistMove caches the transition for the corner orientation coordinate.
// First index is the Twist value (0-2186), second is the Move (0-17).
var TwistMove [2187][18]uint16

// FlipMove caches the transition for the edge orientation coordinate.
// First index is the Flip value (0-2047), second is the Move (0-17).
var FlipMove [2048][18]uint16

// UDSliceMove caches the transition for the UD-slice coordinate.
// First index is the UDSlice value (0-494), second is the Move (0-17).
var UDSliceMove [495][18]uint16

// Phase 2 Move Tables

// CPermMove caches the transition of the corner permutation coordinate (40320 states).
var CPermMove [40320][18]uint16

// EPermMove caches the transition of the edge permutation coordinate (40320 states).
var EPermMove [40320][18]uint16

// MPermMove caches the transition of the middle slice permutation coordinate (24 states).
var MPermMove [24][18]uint16

// MCPermPrun is the pruning table for Phase 2 combining MPerm and CPerm.
// It contains 24 * 40320 states.
var MCPermPrun [24 * 40320]int8

// MEPermPrun is the pruning table for Phase 2 combining MPerm and EPerm.
// It contains 24 * 40320 states.
var MEPermPrun [24 * 40320]int8

// Phase2Moves contains the 10 allowable moves in the G1 group.
// Using any other move would destroy the Phase 1 properties.
var Phase2Moves = []Move{
	MoveU, MoveU2, MoveU3,
	MoveD, MoveD2, MoveD3,
	MoveR2, MoveL2, MoveF2, MoveB2,
}

// InitMoveTables pre-calculates the transition tables for Phase 1.
// It must be called once when the application starts.
func InitMoveTables() {
	// 1. Generate Twist Move Table
	for i := uint16(0); i < 2187; i++ {
		for m := Move(0); m < 18; m++ {
			c := GetPooledCubie() // Take from pool

			c.SetTwist(i)
			c.ApplyMove(m)
			TwistMove[i][m] = c.GetTwist()

			ReleaseCubie(c) // Return to pool!
		}
	}

	// 2. Generate Flip Move Table
	for i := uint16(0); i < 2048; i++ {
		for m := Move(0); m < 18; m++ {
			c := GetPooledCubie()

			c.SetFlip(i)
			c.ApplyMove(m)
			FlipMove[i][m] = c.GetFlip()

			ReleaseCubie(c)
		}
	}

	// 3. Generate UDSlice Move Table
	for i := uint16(0); i < 495; i++ {
		for m := Move(0); m < 18; m++ {
			c := GetPooledCubie()

			c.SetUDSlice(i)
			c.ApplyMove(m)
			UDSliceMove[i][m] = c.GetUDSlice()

			ReleaseCubie(c)
		}
	}

	// 4. Generate CPerm Move Table (40320 states)
	for i := uint16(0); i < 40320; i++ {
		for m := Move(0); m < 18; m++ {
			c := GetPooledCubie()
			c.SetCPerm(i)
			c.ApplyMove(m)
			CPermMove[i][m] = c.GetCPerm()
			ReleaseCubie(c)
		}
	}

	// 5. Generate EPerm Move Table (40320 states)
	for i := uint16(0); i < 40320; i++ {
		for m := Move(0); m < 18; m++ {
			c := GetPooledCubie()
			c.SetEPerm(i)
			c.ApplyMove(m)
			EPermMove[i][m] = c.GetEPerm()
			ReleaseCubie(c)
		}
	}

	// 6. Generate MPerm Move Table (24 states)
	for i := uint16(0); i < 24; i++ {
		for m := Move(0); m < 18; m++ {
			c := GetPooledCubie()
			c.SetMPerm(i)
			c.ApplyMove(m)
			MPermMove[i][m] = c.GetMPerm()
			ReleaseCubie(c)
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
		index := 0 // We track the 1D index manually
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
				index++
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

	// 4. Initialize Phase 2 tables to -1
	for i := 0; i < len(MCPermPrun); i++ {
		MCPermPrun[i] = -1
		MEPermPrun[i] = -1
	}

	// 5. Generate MCPermPrun table
	MCPermPrun[0] = 0
	depth = int8(0)
	done = 1

	for done < len(MCPermPrun) {
		for i := 0; i < len(MCPermPrun); i++ {
			if MCPermPrun[i] == depth {
				mPerm := uint16(i / 40320)
				cPerm := uint16(i % 40320)

				// ONLY use Phase 2 moves!
				for _, m := range Phase2Moves {
					newMPerm := MPermMove[mPerm][m]
					newCPerm := CPermMove[cPerm][m]
					newIndex := int(newMPerm)*40320 + int(newCPerm)

					if MCPermPrun[newIndex] == -1 {
						MCPermPrun[newIndex] = depth + 1
						done++
					}
				}
			}
		}
		depth++
	}

	// 6. Generate MEPermPrun table
	MEPermPrun[0] = 0
	depth = int8(0)
	done = 1

	for done < len(MEPermPrun) {
		for i := 0; i < len(MEPermPrun); i++ {
			if MEPermPrun[i] == depth {
				mPerm := uint16(i / 40320)
				ePerm := uint16(i % 40320)

				// ONLY use Phase 2 moves!
				for _, m := range Phase2Moves {
					newMPerm := MPermMove[mPerm][m]
					newEPerm := EPermMove[ePerm][m]
					newIndex := int(newMPerm)*40320 + int(newEPerm)

					if MEPermPrun[newIndex] == -1 {
						MEPermPrun[newIndex] = depth + 1
						done++
					}
				}
			}
		}
		depth++
	}
}
