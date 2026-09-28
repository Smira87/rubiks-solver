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

// MoveCubieU returns the mathematical representation of the Up move.
func MoveCubieU() *CubieCube {
	return basicCubieMoves[0]
}
