package cube

// TwistMove caches the transition for the corner orientation coordinate.
// First index is the Twist value (0-2186), second is the Move (0-17).
var TwistMove [2187][18]uint16

// FlipMove caches the transition for the edge orientation coordinate.
// First index is the Flip value (0-2047), second is the Move (0-17).
var FlipMove [2048][18]uint16

// UDSliceMove caches the transition for the UD-slice coordinate.
// First index is the UDSlice value (0-494), second is the Move (0-17).
var UDSliceMove [495][18]uint16

// InitMoveTables pre-calculates the transition tables for Phase 1.
// It must be called once when the application starts.

// InitMoveTables pre-calculates the transition tables for Phase 1.
// It must be called once when the application starts.
func InitMoveTables() {
	// 1. Generate Twist Move Table
	for i := uint16(0); i < 2187; i++ {
		for m := Move(0); m < 18; m++ {
			c := NewSolvedCubieCube()
			c.SetTwist(i)
			c.ApplyMove(m)
			TwistMove[i][m] = c.GetTwist()
		}
	}

	// 2. Generate Flip Move Table
	for i := uint16(0); i < 2048; i++ {
		for m := Move(0); m < 18; m++ {
			c := NewSolvedCubieCube()
			c.SetFlip(i)
			c.ApplyMove(m)
			FlipMove[i][m] = c.GetFlip()
		}
	}

	// 3. Generate UDSlice Move Table
	for i := uint16(0); i < 495; i++ {
		for m := Move(0); m < 18; m++ {
			c := NewSolvedCubieCube()
			c.SetUDSlice(i)
			c.ApplyMove(m)
			UDSliceMove[i][m] = c.GetUDSlice()
		}
	}
}
