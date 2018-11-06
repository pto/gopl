package main

import (
	"os"
	"testing"
)

// Empty test to keep the testing package in the imports.
func Test_countLines(t *testing.T) {
}

func Example_noDupes() {
	os.Args = []string{"ex04", "testfile1"}
	main()
	// Output:
}

func Example_oneFile() {
	os.Args = []string{"ex04", "testfile5"}
	main()
	// Unordered output:
	// 2	2nd (testfile5)
	// 3	3rd (testfile5)
}

func Example_twoFiles() {
	os.Args = []string{"ex04", "testfile2", "testfile1"}
	main()
	// Output:
	// 2	This is a test (testfile1, testfile2)
}

func Example_threeFiles() {
	os.Args = []string{"ex04", "testfile1", "testfile2", "testfile3"}
	main()
	// Unordered output:
	// 2	  (testfile3)
	// 2	This is a test (testfile1, testfile2)
	// 2	This is only a test (testfile1, testfile3)
}
