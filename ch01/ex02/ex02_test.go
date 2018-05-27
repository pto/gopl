package main

import "os"

func Example() {
	os.Args = []string{"ex02", "Still", "testing"}
	main()
	// Output:
	// 1 Still
	// 2 testing
}
