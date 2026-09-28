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
