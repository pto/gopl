// Package ex03 has multiple implementations of bit counting.
package ex03

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
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

func PopCountLoop(x uint64) int {
	count := 0
	for i := uint(0); i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

func PopCountShift(x uint64) int {
	count := 0
	for x > 0 {
		count += int(x & 1)
		x >>= 1
	}
	return count
}

func PopCountMask(x uint64) int {
	count := 0
	for x > 0 {
		count++
		x &= x - 1
	}
	return count
}
