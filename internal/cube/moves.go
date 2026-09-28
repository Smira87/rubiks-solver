// Package cube provides the domain model and core mechanics for the Rubik's Cube.
// It implements the facelet representation required for the Kociemba two-phase algorithm.
package cube

// Move represents one of the 18 possible moves in the Kociemba algorithm.
type Move int

// Kociemba algorithm uses moves encoded as integers from 0 to 17.
const (
	MoveU Move = iota
	MoveU2
	MoveU3 // U3 is U'
	MoveR
	MoveR2
	MoveR3 // R3 is R'
	MoveF
	MoveF2
	MoveF3 // F3 is F'
	MoveD
	MoveD2
	MoveD3 // D3 is D'
	MoveL
	MoveL2
	MoveL3 // L3 is L'
	MoveB
	MoveB2
	MoveB3 // B3 is B'
)

// baseMoves contains the 6 foundational permutations (clockwise).
// Carefully mapped according to standard Kociemba facelet layout.
var baseMoves = [6][54]byte{
	// U (Up)
	{
		6, 3, 0, 7, 4, 1, 8, 5, 2, // U
		45, 46, 47, 12, 13, 14, 15, 16, 17, // R
		9, 10, 11, 21, 22, 23, 24, 25, 26, // F
		27, 28, 29, 30, 31, 32, 33, 34, 35, // D
		18, 19, 20, 39, 40, 41, 42, 43, 44, // L
		36, 37, 38, 48, 49, 50, 51, 52, 53, // B
	},
	// R (Right)
	{
		0, 1, 20, 3, 4, 23, 6, 7, 26, // U
		15, 12, 9, 16, 13, 10, 17, 14, 11, // R
		18, 19, 29, 21, 22, 32, 24, 25, 35, // F
		27, 28, 51, 30, 31, 48, 33, 34, 45, // D
		36, 37, 38, 39, 40, 41, 42, 43, 44, // L
		8, 46, 47, 5, 49, 50, 2, 52, 53, // B
	},
	// F (Front)
	{
		0, 1, 2, 3, 4, 5, 44, 41, 38, // U
		6, 10, 11, 7, 13, 14, 8, 16, 17, // R
		24, 21, 18, 25, 22, 19, 26, 23, 20, // F
		15, 12, 9, 30, 31, 32, 33, 34, 35, // D
		36, 37, 27, 39, 40, 28, 42, 43, 29, // L
		45, 46, 47, 48, 49, 50, 51, 52, 53, // B
	},
	// D (Down)
	{
		0, 1, 2, 3, 4, 5, 6, 7, 8, // U
		9, 10, 11, 12, 13, 14, 24, 25, 26, // R
		18, 19, 20, 21, 22, 23, 42, 43, 44, // F
		33, 30, 27, 34, 31, 28, 35, 32, 29, // D
		36, 37, 38, 39, 40, 41, 51, 52, 53, // L
		45, 46, 47, 48, 49, 50, 15, 16, 17, // B
	},
	// L (Left)
	{
		53, 1, 2, 50, 4, 5, 47, 7, 8, // U
		9, 10, 11, 12, 13, 14, 15, 16, 17, // R
		0, 19, 20, 3, 22, 23, 6, 25, 26, // F
		18, 28, 29, 21, 31, 32, 24, 34, 35, // D
		42, 39, 36, 43, 40, 37, 44, 41, 38, // L
		45, 46, 33, 48, 49, 30, 51, 52, 27, // B
	},
	// B (Back)
	{
		11, 14, 17, 3, 4, 5, 6, 7, 8, // U
		9, 10, 35, 12, 13, 34, 15, 16, 33, // R
		18, 19, 20, 21, 22, 23, 24, 25, 26, // F
		27, 28, 29, 30, 31, 32, 36, 39, 42, // D
		2, 37, 38, 1, 40, 41, 0, 43, 44, // L
		51, 48, 45, 52, 49, 46, 53, 50, 47, // B
	},
}

// AllMoves contains all 18 permutations. It is populated dynamically on startup.
var AllMoves [18][54]byte

// init is a special Go function that runs automatically when the package loads.
// We use it to pre-calculate double (e.g., U2) and inverted (e.g., U') moves.
func init() {
	for i := 0; i < 6; i++ {
		move1 := baseMoves[i]

		// Move 1: Clockwise (e.g., U)
		AllMoves[i*3] = move1

		// Move 2: Double turn (e.g., U2) -> Apply move1 to move1
		var move2 [54]byte
		for j := 0; j < 54; j++ {
			move2[j] = move1[move1[j]]
		}
		AllMoves[i*3+1] = move2

		// Move 3: Counter-clockwise (e.g., U') -> Apply move1 to move2
		var move3 [54]byte
		for j := 0; j < 54; j++ {
			move3[j] = move1[move2[j]]
		}
		AllMoves[i*3+2] = move3
	}
}

// ApplyMove applies one of the 18 standard Kociemba moves to the cube.
func (c *Cube) ApplyMove(m Move) {
	perm := AllMoves[m]
	var next [54]Color
	for i := 0; i < 54; i++ {
		next[i] = c.Facelets[perm[i]]
	}
	c.Facelets = next
}
