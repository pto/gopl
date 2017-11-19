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

// PopCountLoop1 returns the population count of x, using a loop (version 1).
func PopCountLoop1(x uint64) int {
	count := 0
	for i := uint(0); i < 8; i++ {
		count += int(pc[byte(x>>(i*8))])
	}
	return count
}

// PopCountLoop1 returns the population count of x, using a loop (version 2).
func PopCountLoop2(x uint64) int {
	count := 0
	for i := 0; i < 8; i++ {
		count += int(pc[byte(x)])
		x >>= 8
	}
	return count
}

// PopCountShift returns the population count of x, shifting each bit.
func PopCountShift(x uint64) int {
	count := 0
	for x > 0 {
		count += int(x & 1)
		x >>= 1
	}
	return count
}

// PopCountMask returns the population count of x, clearing the rightmost
// non-zero bit.
func PopCountMask(x uint64) int {
	count := 0
	for x > 0 {
		count++
		x &= x - 1
	}
	return count
}
