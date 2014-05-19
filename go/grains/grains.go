/*
Package grains implements:

Square; a function that takes an integer
representing a square postion on a chessboard.
Returns a grain count such that grains = 2^(number - 1).

Total; a function that returns the total number
of grains on a chess board, such that total grains = 2^(64-1)
*/
package grains

func Square(number int) uint64 {
	return 1 << uint64(number-1)
}

func Total() uint64 {
	return (1 << 64) - 1
}
