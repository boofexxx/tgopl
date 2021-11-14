package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountClear(x uint64) int {
	t := 0
	for x != 0 {
		x = x & (x - 1)
		t++
	}
	return t
}

func PopCountShift(x uint64) int {
	t := 0
	for i := 0; i < 64; i++ {
		t += int((x >> i) & 1)
	}
	return t
}

func PopCountIter(x uint64) int {
	t := 0
	for i := 0; i < 8; i++ {
		t += int(pc[byte(x>>(i*8))])
	}
	return t
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
