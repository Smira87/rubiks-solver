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
