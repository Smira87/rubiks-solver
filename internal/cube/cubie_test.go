package cube_test

import (
	"testing"

	"rubiks-solver/internal/cube"
)

func TestNewSolvedCubieCube(t *testing.T) {
	// Arrange
	cc := cube.NewSolvedCubieCube()

	// Assert Corners
	// In a solved cube, the URF (Up-Right-Front) corner should be at index 0
	// and its orientation must be 0.
	if cc.CP[cube.URF] != cube.URF || cc.CO[cube.URF] != 0 {
		t.Errorf("Expected URF corner to be solved (pos URF, ori 0), got pos %v, ori %v",
			cc.CP[cube.URF], cc.CO[cube.URF])
	}

	// Assert Edges
	// The UR (Up-Right) edge should be at index 0 with orientation 0.
	if cc.EP[cube.UR] != cube.UR || cc.EO[cube.UR] != 0 {
		t.Errorf("Expected UR edge to be solved (pos UR, ori 0), got pos %v, ori %v",
			cc.EP[cube.UR], cc.EO[cube.UR])
	}
}

func TestCubieCube_Multiply_SolvedBySolved(t *testing.T) {
	// Arrange: Two completely solved cubes
	c1 := cube.NewSolvedCubieCube()
	c2 := cube.NewSolvedCubieCube()

	// Act: Multiply c1 by c2 (c1 will store the result)
	c1.Multiply(c2)

	// Assert: It should still be perfectly solved
	solved := cube.NewSolvedCubieCube()

	for i := 0; i < 8; i++ {
		if c1.CP[i] != solved.CP[i] || c1.CO[i] != solved.CO[i] {
			t.Errorf("Corner %d changed after multiplying solved cubes", i)
		}
	}
	for i := 0; i < 12; i++ {
		if c1.EP[i] != solved.EP[i] || c1.EO[i] != solved.EO[i] {
			t.Errorf("Edge %d changed after multiplying solved cubes", i)
		}
	}
}

func TestCubieCube_Multiply_U4TimesReturnsToSolved(t *testing.T) {
	// Arrange
	c := cube.NewSolvedCubieCube()

	// We need the predefined U move from our basicCubieMoves.
	// Note: We'll export a function later to access these, but for tests
	// inside the same package, we can access package-level vars directly.

	// Act: Apply U move 4 times using Multiply
	c.Multiply(cube.MoveCubieU())
	c.Multiply(cube.MoveCubieU())
	c.Multiply(cube.MoveCubieU())
	c.Multiply(cube.MoveCubieU())

	// Assert
	solved := cube.NewSolvedCubieCube()
	if c.CP != solved.CP || c.EP != solved.EP {
		t.Errorf("CubieCube did not return to solved state after U4")
	}
}
