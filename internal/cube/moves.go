// Package cube provides the domain model and core mechanics for the Rubik's Cube.
// It implements the facelet representation required for the Kociemba two-phase algorithm.
package cube

// moveU is the permutation array for the Up face (90-degree clockwise).
// The value at index 'i' tells us which facelet moves TO position 'i'.
var moveU = [54]byte{
	6, 3, 0, 7, 4, 1, 8, 5, 2, // U face rotated
	45, 46, 47, 12, 13, 14, 15, 16, 17, // R face (top row gets B)
	9, 10, 11, 21, 22, 23, 24, 25, 26, // F face (top row gets R)
	27, 28, 29, 30, 31, 32, 33, 34, 35, // D face (unchanged)
	18, 19, 20, 39, 40, 41, 42, 43, 44, // L face (top row gets F)
	36, 37, 38, 48, 49, 50, 51, 52, 53, // B face (top row gets L)
}

// moveR is the permutation array for the Right face.
var moveR = [54]byte{
	0, 1, 20, 3, 4, 23, 6, 7, 26, // U gets F
	15, 12, 9, 16, 13, 10, 17, 14, 11, // R rotated
	18, 19, 29, 21, 22, 32, 24, 25, 35, // F gets D
	27, 28, 51, 30, 31, 48, 33, 34, 45, // D gets B (reversed orientation)
	36, 37, 38, 39, 40, 41, 42, 43, 44, // L unchanged
	8, 46, 47, 5, 49, 50, 2, 52, 53, // B gets U (reversed orientation)
}

// ApplyPermutation updates the cube's state using a given permutation array.
// This is the core engine for all Kociemba movements.
func (c *Cube) ApplyPermutation(perm [54]byte) {
	var next [54]Color
	for i := 0; i < 54; i++ {
		next[i] = c.Facelets[perm[i]]
	}
	c.Facelets = next
}

// MoveU applies the U move.
func (c *Cube) MoveU() {
	c.ApplyPermutation(moveU)
}

// MoveR applies the R move.
func (c *Cube) MoveR() {
	c.ApplyPermutation(moveR)
}
