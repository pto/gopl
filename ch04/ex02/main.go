// Ex02 prints the hash of its standard input, defaulting to SHA256, with
// command line flags for SHA384 and SHA512.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

var useSHA384 = flag.Bool("sha384", false, "use SHA384 instead of SHA256")
var useSHA512 = flag.Bool("sha512", false, "use SHA512 instead of SHA256")

func main() {
	flag.Parse()
	if *useSHA384 && *useSHA512 {
		fmt.Fprintln(os.Stderr, "ex02: cannot specify multiple flags")
		flag.PrintDefaults()
		os.Exit(1)
	}

	var h hash.Hash

	switch {
	case *useSHA384:
		h = sha512.New384()
	case *useSHA512:
		h = sha512.New()
	default:
		h = sha256.New()
	}
	_, err := io.Copy(h, os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ex02:", err)
		os.Exit(1)
	}
	fmt.Printf("%x\n", h.Sum(nil))
}
