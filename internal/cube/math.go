package cube

// Cnk calculates the binomial coefficient "n choose k".
// It is heavily used in Kociemba's algorithm to calculate permutation coordinates.
func Cnk(n, k int) int {
	if n < k {
		return 0
	}
	if k > n/2 {
		k = n - k // Symmetry property: C(n, k) == C(n, n-k)
	}

	res := 1
	for i := 1; i <= k; i++ {
		res *= n - i + 1
		res /= i
	}

	return res
}
