package util

// PowU8 returns in^pow
func PowU8(in, pow uint8) (val uint8) {
	val = 1

	for ; pow > 0; pow-- {
		val *= in
	}

	return
}
