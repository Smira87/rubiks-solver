//go:build js && wasm

package main

import (
	"strings"
	"syscall/js"

	"rubiks-solver/internal/cube"
	"rubiks-solver/internal/solver"
)

// moveNames maps our internal integer moves back to standard Rubik's Cube notation.
var moveNames = []string{
	"U", "U2", "U'", "R", "R2", "R'", "F", "F2", "F'",
	"D", "D2", "D'", "L", "L2", "L'", "B", "B2", "B'",
}

// parseScramble converts a string like "R U F2" into a slice of Moves.
func parseScramble(scramble string) []cube.Move {
	// Create a map for fast lookup
	strToMove := make(map[string]cube.Move)
	for i, name := range moveNames {
		strToMove[name] = cube.Move(i)
	}

	var moves []cube.Move
	parts := strings.Fields(scramble) // Splits by space
	for _, p := range parts {
		if m, exists := strToMove[p]; exists {
			moves = append(moves, m)
		}
	}
	return moves
}

// solveCubeWrapper is the function exposed to JavaScript.
// JS will call it like this: solver.solve("R U F2")
func solveCubeWrapper(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		return "Error: No scramble provided"
	}

	scrambleStr := args[0].String()
	scrambleMoves := parseScramble(scrambleStr)

	// Create a new cube and apply the scramble
	c := cube.NewSolvedCubieCube()
	for _, m := range scrambleMoves {
		c.ApplyMove(m)
	}

	// Solve the cube (max 22 moves)
	solution := solver.Solve(c, 22)

	if solution == nil {
		return "Error: Unsolvable state or algorithm failed"
	}

	// Convert solution moves back to a string
	var result []string
	for _, m := range solution {
		result = append(result, moveNames[m])
	}

	return strings.Join(result, " ")
}

func main() {
	// 1. Initialize tables on startup.
	// This takes ~0.4s and runs when the browser loads the page.
	cube.InitMoveTables()
	cube.InitPruningTables()

	// 2. Bind our Go function to the global JavaScript 'window' object.
	js.Global().Set("kociembaSolve", js.FuncOf(solveCubeWrapper))

	// 3. Prevent the Go program from exiting.
	// In WASM, if main() finishes, the functions are destroyed.
	<-make(chan struct{})
}
