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
