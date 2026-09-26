package cube

// Color represents the color of a facelet.
// We use byte for memory efficiency.
type Color byte

// Define constants for each face (using iota for auto-numbering 0, 1, 2...)
const (
	U Color = iota // Up
	R              // Right
	F              // Front
	D              // Down
	L              // Left
	B              // Back
)

// Cube represents the state of a Rubik's Cube (54 facelets)
type Cube struct {
	Facelets [54]Color
}

// NewSolvedCube creates and returns a perfectly solved cube
func NewSolvedCube() *Cube {
	var facelets [54]Color

	// Initialize the solved state
	// Kociemba order: U (0-8), R (9-17), F (18-26), D (27-35), L (36-44), B (45-53)
	for i := 0; i < 54; i++ {
		// Dividing by 9 gives us the face number (0 for U, 1 for R, etc.)
		facelets[i] = Color(i / 9)
	}

	return &Cube{
		Facelets: facelets,
	}
}

func (c *Cube) MoveU() {
	// Create a copy of the current facelets state.
	// We need this to read old values while overwriting the array.
	old := c.Facelets

	// 1. Rotate the U face itself (indices 0 to 8)
	// Corners
	c.Facelets[0] = old[6] // UBL gets UFL
	c.Facelets[2] = old[0] // UBR gets UBL
	c.Facelets[8] = old[2] // UFR gets UBR
	c.Facelets[6] = old[8] // UFL gets UFR
	// Edges
	c.Facelets[1] = old[3] // UB gets UL
	c.Facelets[5] = old[1] // UR gets UB
	c.Facelets[7] = old[5] // UF gets UR
	c.Facelets[3] = old[7] // UL gets UF
	// Center (index 4) remains unchanged

	// 2. Rotate the adjacent top rows of R, F, L, B faces
	// Right face top row (9, 10, 11) gets Back face top row (45, 46, 47)
	c.Facelets[9] = old[45]
	c.Facelets[10] = old[46]
	c.Facelets[11] = old[47]

	// Front face top row (18, 19, 20) gets Right face top row (9, 10, 11)
	c.Facelets[18] = old[9]
	c.Facelets[19] = old[10]
	c.Facelets[20] = old[11]

	// Left face top row (36, 37, 38) gets Front face top row (18, 19, 20)
	c.Facelets[36] = old[18]
	c.Facelets[37] = old[19]
	c.Facelets[38] = old[20]

	// Back face top row (45, 46, 47) gets Left face top row (36, 37, 38)
	c.Facelets[45] = old[36]
	c.Facelets[46] = old[37]
	c.Facelets[47] = old[38]
}
