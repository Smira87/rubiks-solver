package cube_test

import (
	"testing"

	"rubiks-solver/internal/cube"
)

func TestCubieCube_Twist(t *testing.T) {
	// Arrange
	c := cube.NewSolvedCubieCube()

	// 1. A solved cube should have a Twist coordinate of 0
	if twist := c.GetTwist(); twist != 0 {
		t.Errorf("Expected Twist 0 for a solved cube, got %d", twist)
	}

	// 2. Property-based testing: Set a twist, get it back, and check if it matches.
	// The maximum value for Twist is 3^7 - 1 = 2186.
	for expectedTwist := uint16(0); expectedTwist < 2187; expectedTwist++ {
		c.SetTwist(expectedTwist)

		actualTwist := c.GetTwist()
		if actualTwist != expectedTwist {
			t.Errorf("SetTwist(%d) but GetTwist() returned %d", expectedTwist, actualTwist)
		}
	}
}

func TestCubieCube_Flip(t *testing.T) {
	// Arrange
	c := cube.NewSolvedCubieCube()

	// 1. A solved cube should have a Flip coordinate of 0
	if flip := c.GetFlip(); flip != 0 {
		t.Errorf("Expected Flip 0 for a solved cube, got %d", flip)
	}

	// 2. Property-based testing: Set a flip, get it back, and check if it matches.
	// The maximum value for Flip is 2^11 - 1 = 2047.
	for expectedFlip := uint16(0); expectedFlip < 2048; expectedFlip++ {
		c.SetFlip(expectedFlip)

		actualFlip := c.GetFlip()
		if actualFlip != expectedFlip {
			t.Errorf("SetFlip(%d) but GetFlip() returned %d", expectedFlip, actualFlip)
		}
	}
}

func TestCubieCube_UDSlice(t *testing.T) {
	// Arrange
	c := cube.NewSolvedCubieCube()

	// 1. A solved cube should have a UDSlice coordinate of 0
	if slice := c.GetUDSlice(); slice != 0 {
		t.Errorf("Expected UDSlice 0 for a solved cube, got %d", slice)
	}

	// 2. Property-based testing: Set a slice coordinate, get it back, and check if it matches.
	// The maximum value for UDSlice is C(12, 4) - 1 = 494.
	for expectedSlice := uint16(0); expectedSlice < 495; expectedSlice++ {
		c.SetUDSlice(expectedSlice)

		actualSlice := c.GetUDSlice()
		if actualSlice != expectedSlice {
			t.Errorf("SetUDSlice(%d) but GetUDSlice() returned %d", expectedSlice, actualSlice)
		}
	}
}
