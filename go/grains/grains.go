// Package grains implements functions to count the
// number of grains on a chessboard.
package grains

// Square takes an integer representing a square postion on a chessboard.
// Returns a grain count such that grains = 2^(number - 1).
func Square(number int) uint64 {
	return 1 << uint64(number-1)
}

// Total returns the total number of grains on a chessboard
// such that total grains = 2^(64-1).
func Total() uint64 {
	return (1 << 64) - 1
}
