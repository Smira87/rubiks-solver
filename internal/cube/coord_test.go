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

func TestCubieCube_CPerm(t *testing.T) {
	// Arrange
	c := cube.NewSolvedCubieCube()

	// 1. A solved cube should have a CPerm coordinate of 0
	if cperm := c.GetCPerm(); cperm != 0 {
		t.Errorf("Expected CPerm 0 for a solved cube, got %d", cperm)
	}

	// 2. Property-based testing: Set a permutation, get it back, and check if it matches.
	// The maximum value for Corner Permutation is 8! - 1 = 40319.
	for expectedCPerm := uint16(0); expectedCPerm < 40320; expectedCPerm++ {
		c.SetCPerm(expectedCPerm)

		actualCPerm := c.GetCPerm()
		if actualCPerm != expectedCPerm {
			t.Errorf("SetCPerm(%d) but GetCPerm() returned %d", expectedCPerm, actualCPerm)
		}
	}
}

func TestCubieCube_EPerm(t *testing.T) {
	// Arrange
	c := cube.NewSolvedCubieCube()

	// 1. A solved cube should have an EPerm coordinate of 0
	if eperm := c.GetEPerm(); eperm != 0 {
		t.Errorf("Expected EPerm 0 for a solved cube, got %d", eperm)
	}

	// 2. Property-based testing: Set an EPerm coordinate, get it back, and check if it matches.
	// The maximum value for EPerm is 8! - 1 = 40319.
	// This covers the permutations of the 8 U/D layer edges.
	for expectedEPerm := uint16(0); expectedEPerm < 40320; expectedEPerm++ {
		c.SetEPerm(expectedEPerm)

		actualEPerm := c.GetEPerm()
		if actualEPerm != expectedEPerm {
			t.Errorf("SetEPerm(%d) but GetEPerm() returned %d", expectedEPerm, actualEPerm)
		}
	}
}

func TestCubieCube_MPerm(t *testing.T) {
	// Arrange
	c := cube.NewSolvedCubieCube()

	// 1. A solved cube should have an MPerm coordinate of 0
	if mperm := c.GetMPerm(); mperm != 0 {
		t.Errorf("Expected MPerm 0 for a solved cube, got %d", mperm)
	}

	// 2. Property-based testing: Set an MPerm coordinate, get it back.
	// The maximum value for MPerm is 4! - 1 = 23.
	// This covers the permutations of the 4 equatorial slice edges (FR, FL, BL, BR).
	for expectedMPerm := uint16(0); expectedMPerm < 24; expectedMPerm++ {
		c.SetMPerm(expectedMPerm)

		actualMPerm := c.GetMPerm()
		if actualMPerm != expectedMPerm {
			t.Errorf("SetMPerm(%d) but GetMPerm() returned %d", expectedMPerm, actualMPerm)
		}
	}
}
