package grains

func Square(number int) uint64 {
	return uint64(1 << uint64(number-1))
}

func Total() uint64 {
	return uint64(1<<63) + (uint64(1<<63) - 1)
}
