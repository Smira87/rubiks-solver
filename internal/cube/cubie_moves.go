package cube

// basicCubieMoves contains the 6 foundational Kociemba moves represented as CubieCubes.
// These are mathematically pre-calculated to alter permutations and orientations correctly.
var basicCubieMoves = [6]*CubieCube{
	// U (Up)
	{
		CP: [8]Corner{UBR, URF, UFL, ULB, DFR, DLF, DBL, DRB},
		CO: [8]byte{0, 0, 0, 0, 0, 0, 0, 0},
		EP: [12]Edge{UB, UR, UF, UL, DR, DF, DL, DB, FR, FL, BL, BR},
		EO: [12]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	// R (Right)
	{
		CP: [8]Corner{DFR, UFL, ULB, URF, DRB, DLF, DBL, UBR},
		CO: [8]byte{2, 0, 0, 1, 1, 0, 0, 2},
		EP: [12]Edge{FR, UF, UL, UB, BR, DF, DL, DB, DR, FL, BL, UR},
		EO: [12]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	// F (Front)
	{
		CP: [8]Corner{UFL, DLF, ULB, UBR, URF, DFR, DBL, DRB},
		CO: [8]byte{1, 2, 0, 0, 2, 1, 0, 0},
		EP: [12]Edge{UR, FL, UL, UB, DR, FR, DL, DB, UF, DF, BL, BR},
		EO: [12]byte{0, 1, 0, 0, 0, 1, 0, 0, 1, 1, 0, 0},
	},
	// D (Down)
	{
		CP: [8]Corner{URF, UFL, ULB, UBR, DLF, DBL, DRB, DFR},
		CO: [8]byte{0, 0, 0, 0, 0, 0, 0, 0},
		EP: [12]Edge{UR, UF, UL, UB, DF, DL, DB, DR, FR, FL, BL, BR},
		EO: [12]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	// L (Left)
	{
		CP: [8]Corner{URF, ULB, DBL, UBR, DFR, UFL, DLF, DRB},
		CO: [8]byte{0, 1, 2, 0, 0, 2, 1, 0},
		EP: [12]Edge{UR, UF, BL, UB, DR, DF, FL, DB, FR, UL, DL, BR},
		EO: [12]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	},
	// B (Back)
	{
		CP: [8]Corner{URF, UFL, UBR, DRB, DFR, DLF, ULB, DBL},
		CO: [8]byte{0, 0, 1, 2, 0, 0, 2, 1},
		EP: [12]Edge{UR, UF, UL, BR, DR, DF, DL, BL, FR, FL, UB, DB},
		EO: [12]byte{0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 1},
	},
}

// AllCubieMoves holds all 18 standard Kociemba moves at the cubie level.
// It is dynamically populated on startup to save memory and avoid hardcoding.
var AllCubieMoves [18]*CubieCube

// init runs automatically when the cube package is loaded.
// We use it to pre-calculate double (e.g., U2) and inverted (e.g., U') moves using math.
func init() {
	for i := 0; i < 6; i++ {
		// Base move (clockwise, e.g., U, R, F, D, L, B)
		move1 := basicCubieMoves[i]
		AllCubieMoves[i*3] = move1

		// Double move (e.g., U2) -> Multiply move1 by move1
		move2 := NewSolvedCubieCube()
		move2.Multiply(move1)
		move2.Multiply(move1)
		AllCubieMoves[i*3+1] = move2

		// Inverse move (e.g., U') -> Multiply move2 by move1
		move3 := NewSolvedCubieCube()
		move3.Multiply(move2)
		move3.Multiply(move1)
		AllCubieMoves[i*3+2] = move3
	}
}

// ApplyMove applies one of the 18 standard Kociemba moves to the CubieCube.
// This is done by multiplying the current state with the pre-calculated move state.
func (c *CubieCube) ApplyMove(m Move) {
	c.Multiply(AllCubieMoves[m])
}
