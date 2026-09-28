package cube

// Corner represents the 8 corners of the cube.
type Corner byte

const (
	URF Corner = iota // Up-Right-Front
	UFL               // Up-Front-Left
	ULB               // Up-Left-Back
	UBR               // Up-Back-Right
	DFR               // Down-Front-Right
	DLF               // Down-Left-Front
	DBL               // Down-Back-Left
	DRB               // Down-Right-Back
)

// Edge represents the 12 edges of the cube.
type Edge byte

const (
	UR Edge = iota // Up-Right
	UF             // Up-Front
	UL             // Up-Left
	UB             // Up-Back
	DR             // Down-Right
	DF             // Down-Front
	DL             // Down-Left
	DB             // Down-Back
	FR             // Front-Right
	FL             // Front-Left
	BL             // Back-Left
	BR             // Back-Right
)

// CubieCube represents the cube on the "Cubie Level" (corners and edges).
// This is the essential bridge between the UI (Facelets) and the solver (Coordinates).
type CubieCube struct {
	// CP (Corner Permutation) stores which corner piece is at a given corner position.
	CP [8]Corner
	// CO (Corner Orientation) stores the twist of the corner piece (0, 1, or 2).
	CO [8]byte
	// EP (Edge Permutation) stores which edge piece is at a given edge position.
	EP [12]Edge
	// EO (Edge Orientation) stores the flip of the edge piece (0 or 1).
	EO [12]byte
}

// NewSolvedCubieCube creates a CubieCube in the completely solved state.
func NewSolvedCubieCube() *CubieCube {
	cc := &CubieCube{}

	for i := 0; i < 8; i++ {
		cc.CP[i] = Corner(i)
		cc.CO[i] = 0 // 0 means correctly oriented
	}

	for i := 0; i < 12; i++ {
		cc.EP[i] = Edge(i)
		cc.EO[i] = 0 // 0 means correctly oriented
	}

	return cc
}

// Multiply multiplies the current CubieCube (c) with another CubieCube (b).
// In Kociemba's math, applying a move is multiplying the current state by a move state.
func (c *CubieCube) Multiply(b *CubieCube) {
	// Temporary arrays to hold the new state during calculation
	var newCP [8]Corner
	var newCO [8]byte
	var newEP [12]Edge
	var newEO [12]byte

	// 1. Multiply Corners
	for i := 0; i < 8; i++ {
		// Permutation: apply b's permutation to c
		newCP[i] = c.CP[b.CP[i]]

		// Orientation: add orientations and modulo 3 (since corners have 3 states: 0, 1, 2)
		// We get the old orientation from the corner that moved here, and add the new twist.
		ori := c.CO[b.CP[i]] + b.CO[i]

		// Handle standard edge cases in corner twists (Kociemba specific)
		if ori >= 3 {
			newCO[i] = ori - 3
		} else {
			newCO[i] = ori
		}
	}

	// 2. Multiply Edges
	for i := 0; i < 12; i++ {
		// Permutation: apply b's permutation to c
		newEP[i] = c.EP[b.EP[i]]

		// Orientation: add orientations and modulo 2 (edges have 2 states: 0, 1)
		ori := c.EO[b.EP[i]] + b.EO[i]
		if ori >= 2 {
			newEO[i] = ori - 2
		} else {
			newEO[i] = ori
		}
	}

	// Apply the calculated state back to the current cube
	c.CP = newCP
	c.CO = newCO
	c.EP = newEP
	c.EO = newEO
}
