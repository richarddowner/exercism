package grains

func Square(number int) uint64 {
	return 1 << uint64(number-1)
}

func Total() uint64 {
	return (1 << 64) - 1
}
