// Ex04 creates a chain of goroutines connected by channels and prints
// the time needed to transit the entire pipeline.
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func stage(in <-chan int, out chan<- int) {
	out <- (<-in + 1)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		usage()
	}

	root := make(chan int)
	next := root
	start := time.Now()
	for i := 0; i < n; i++ {
		prev := next
		next = make(chan int)
		go stage(prev, next)
	}
	fmt.Println("Starting goroutines:", time.Since(start))

	start = time.Now()
	root <- 0
	fmt.Println("Result:", <-next)
	fmt.Println("Transiting pipeline:", time.Since(start))
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: ex04 <number of goroutines>")
	os.Exit(1)
}
