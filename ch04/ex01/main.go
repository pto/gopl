// Ex01 counts the number of differing bits in two SHA256 hashes.
package main

import (
	"crypto/sha256"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: <string1> <string2>")
		os.Exit(1)
	}
	s1 := sha256.Sum256([]byte(os.Args[1]))
	s2 := sha256.Sum256([]byte(os.Args[2]))
	diff := s1
	count := 0
	for i, _ := range diff {
		diff[i] = s1[i] ^ s2[i]
		count += bits.OnesCount8(diff[i])
	}
	fmt.Printf("1st:        %x\n", s1)
	fmt.Printf("2nd:        %x\n", s2)
	fmt.Printf("1st ^ 2nd:  %x\n", diff)
	fmt.Printf("Difference: %d bits out of %d\n", count, len(diff)*8)
}
