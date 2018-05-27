package main

import "os"

func Example() {
	os.Args = []string{"ex01", "This", "is", "a", "test"}
	main()
	// Output:
	// ex01 This is a test
}
