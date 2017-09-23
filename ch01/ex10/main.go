// Ex10 fetches URLs in parallel, reports their times and sizes, and saves
// the output in a series of files.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for index, url := range os.Args[1:] {
		go fetch(url, index, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed [%d]\n", time.Since(start).Seconds(), os.Getpid())
}

func fetch(url string, index int, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	filename := fmt.Sprintf("ex10.%d.%d", os.Getpid(), index)
	f, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprintf("while opening %s: %v", filename, err)
		return
	}
	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close() // don't leak resources
	f.Close()
	if err != nil {
		ch <- fmt.Sprintf("while copying %s to %s: %v", url, filename, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
