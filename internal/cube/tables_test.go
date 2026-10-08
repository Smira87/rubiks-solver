package cube_test

import (
	"testing"

	"rubiks-solver/internal/cube"
)

func TestPhase1MoveTables(t *testing.T) {
	// Arrange: Generate the tables
	cube.InitMoveTables()

	// 1. Test Twist Move Table
	// We take a solved cube, apply MoveF, and see what the Twist coordinate becomes.
	c1 := cube.NewSolvedCubieCube()
	c1.ApplyMove(cube.MoveF)
	expectedTwist := c1.GetTwist()

	// The table should have this exact value for Twist=0 and Move=F
	if actual := cube.TwistMove[0][cube.MoveF]; actual != expectedTwist {
		t.Errorf("TwistMove[0][F] expected %d, got %d", expectedTwist, actual)
	}

	// 2. Test Flip Move Table
	c2 := cube.NewSolvedCubieCube()
	c2.ApplyMove(cube.MoveR)
	expectedFlip := c2.GetFlip()

	if actual := cube.FlipMove[0][cube.MoveR]; actual != expectedFlip {
		t.Errorf("FlipMove[0][R] expected %d, got %d", expectedFlip, actual)
	}

	// 3. Test UDSlice Move Table
	c3 := cube.NewSolvedCubieCube()
	c3.ApplyMove(cube.MoveF2)
	expectedUDSlice := c3.GetUDSlice()

	if actual := cube.UDSliceMove[0][cube.MoveF2]; actual != expectedUDSlice {
		t.Errorf("UDSliceMove[0][F2] expected %d, got %d", expectedUDSlice, actual)
	}
}

func TestPhase1PruningTables(t *testing.T) {
	// Arrange: We need move tables generated first, because pruning relies on them.
	cube.InitMoveTables()
	cube.InitPruningTables()

	// 1. Solved state (index 0) must have a distance of 0
	if depth := cube.SliceTwistPrun[0]; depth != 0 {
		t.Errorf("Expected SliceTwist distance 0 for solved state, got %d", depth)
	}
	if depth := cube.SliceFlipPrun[0]; depth != 0 {
		t.Errorf("Expected SliceFlip distance 0 for solved state, got %d", depth)
	}

	// 2. Apply exactly 1 move (e.g., MoveF) to a solved state.
	// The distance must become 1.
	twistAfterF := cube.TwistMove[0][cube.MoveF]
	sliceAfterF := cube.UDSliceMove[0][cube.MoveF]

	// We map two coordinates (Slice and Twist) into a single 1D array index
	indexF := int(sliceAfterF)*2187 + int(twistAfterF)

	if depth := cube.SliceTwistPrun[indexF]; depth != 1 {
		t.Errorf("Expected SliceTwist distance 1 after F move, got %d", depth)
	}
}

func TestPhase2MoveTables(t *testing.T) {
	// Arrange: Initialize tables
	cube.InitMoveTables()

	// 1. Test CPerm Move Table
	c1 := cube.NewSolvedCubieCube()
	c1.ApplyMove(cube.MoveR2) // R2 is a valid Phase 2 move
	expectedCPerm := c1.GetCPerm()

	if actual := cube.CPermMove[0][cube.MoveR2]; actual != expectedCPerm {
		t.Errorf("CPermMove[0][R2] expected %d, got %d", expectedCPerm, actual)
	}

	// 2. Test EPerm Move Table
	c2 := cube.NewSolvedCubieCube()
	c2.ApplyMove(cube.MoveU) // U is a valid Phase 2 move
	expectedEPerm := c2.GetEPerm()

	if actual := cube.EPermMove[0][cube.MoveU]; actual != expectedEPerm {
		t.Errorf("EPermMove[0][U] expected %d, got %d", expectedEPerm, actual)
	}

	// 3. Test MPerm Move Table
	c3 := cube.NewSolvedCubieCube()
	// To affect the middle layer, we need a move like R2, L2, F2, or B2.
	c3.ApplyMove(cube.MoveF2)
	expectedMPerm := c3.GetMPerm()

	if actual := cube.MPermMove[0][cube.MoveF2]; actual != expectedMPerm {
		t.Errorf("MPermMove[0][F2] expected %d, got %d", expectedMPerm, actual)
	}
}
