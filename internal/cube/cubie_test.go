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
