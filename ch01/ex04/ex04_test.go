package main

import "os"

func ExampleNoDupes() {
	os.Args = []string{"ex04", "testfile1"}
	main()
	// Output:
}

func ExampleOneFile() {
	os.Args = []string{"ex04", "testfile5"}
	main()
	// Output (unordered):
	// 2	2nd (testfile5)
	// 3	3rd (testfile5)
}
