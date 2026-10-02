package cube

// GetTwist calculates the corner orientation coordinate (0 to 2186).
// It treats the first 7 corner orientations as a base-3 number.
func (c *CubieCube) GetTwist() uint16 {
	var twist uint16 = 0
	for i := 0; i < 7; i++ {
		// Multiply by 3 (shift base-3) and add the current corner's orientation
		twist = 3*twist + uint16(c.CO[i])
	}
	return twist
}

// SetTwist applies a given corner orientation coordinate (0 to 2186) to the CubieCube.
func (c *CubieCube) SetTwist(twist uint16) {
	var twistParity byte = 0

	// We decode the base-3 number starting from the 6th index down to 0
	for i := 6; i >= 0; i-- {
		c.CO[i] = byte(twist % 3)
		twistParity += c.CO[i]
		twist /= 3
	}

	// The orientation of the 8th corner (index 7) is completely determined by the first 7.
	// The rule of the Rubik's Cube: the sum of all corner twists must be divisible by 3.
	c.CO[7] = (3 - (twistParity % 3)) % 3
}

// GetFlip calculates the edge orientation coordinate (0 to 2047).
// It treats the first 11 edge orientations as a base-2 (binary) number.
func (c *CubieCube) GetFlip() uint16 {
	var flip uint16 = 0
	for i := 0; i < 11; i++ {
		// Bitwise shift left by 1 (same as multiplying by 2)
		// and bitwise OR with the current edge orientation.
		flip = (flip << 1) | uint16(c.EO[i])
	}
	return flip
}

// SetFlip applies a given edge orientation coordinate (0 to 2047) to the CubieCube.
func (c *CubieCube) SetFlip(flip uint16) {
	var flipParity byte = 0

	// We decode the binary number starting from the 10th index down to 0
	for i := 10; i >= 0; i-- {
		// Bitwise AND with 1 gets the lowest bit (same as modulo 2)
		c.EO[i] = byte(flip & 1)

		// XOR operator (^) is perfect for calculating parity (mod 2 addition)
		flipParity ^= c.EO[i]

		// Bitwise shift right by 1 (same as integer division by 2)
		flip >>= 1
	}

	// The orientation of the 12th edge (index 11) ensures total parity is 0.
	c.EO[11] = flipParity
}

// GetUDSlice calculates the UD-slice coordinate (0 to 494).
// It identifies the positions of the 4 equatorial edges (FR, FL, BL, BR).
func (c *CubieCube) GetUDSlice() uint16 {
	var slice uint16 = 0
	k := 3 // We are looking for 4 edges, so index goes from 3 down to 0

	// Scan through all 12 edge positions from BR (11) down to UR (0)
	for j := 11; j >= 0; j-- {
		// Check if the piece at position j is one of the slice edges (FR=8, FL=9, BL=10, BR=11)
		if c.EP[j] >= FR && c.EP[j] <= BR {
			k--
			if k < 0 {
				break
			}
		} else {
			// If it's not a slice edge, we add the combinations to our index
			slice += uint16(Cnk(j, k))
		}
	}

	return slice
}

// SetUDSlice applies a given UD-slice coordinate (0 to 494) to the CubieCube.
// Note: This only places the 4 slice edges correctly. The other 8 edges are filled with dummy values (UR).
// This is perfectly fine because Phase 1 pruning tables only care about the slice edges.
func (c *CubieCube) SetUDSlice(slice uint16) {
	// First, fill all edges with a dummy value (e.g., UR)
	for i := 0; i < 12; i++ {
		c.EP[i] = UR
	}

	x := int(slice)
	k := 3

	for j := 11; j >= 0; j-- {
		val := Cnk(j, k)
		if x >= val {
			// Not a slice edge
			x -= val
		} else {
			// It is a slice edge. Place FR, FL, BL, BR (which are enum values 8, 9, 10, 11).
			// We place them in order: 8 + k.
			c.EP[j] = Edge(8 + k)
			k--
			if k < 0 {
				break
			}
		}
	}
}
