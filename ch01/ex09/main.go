// Ex09 outputs the content found at a URL and the HTTP status code.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var verbose = flag.Bool("V", false, "verbose")

func main() {
	flag.Parse()
	for _, url := range flag.Args() {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex09: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex09: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		if !*verbose {
			fmt.Println(b)
		}
		fmt.Println("Status:", resp.Status)
		if *verbose {
			fmt.Println("Proto:", resp.Proto)
			for k, v := range resp.Header {
				fmt.Printf("Header[%q]: %v\n", k, strings.Join(v, ", "))
			}
			fmt.Println("ContentLength:", resp.ContentLength)
			fmt.Println("TransferEncoding:",
				strings.Join(resp.TransferEncoding, ", "))
			fmt.Println("Close:", resp.Close)
			for k, v := range resp.Trailer {
				fmt.Printf("Trailer[%q]: %v\n", k, strings.Join(v, ", "))
			}
			if resp.TLS != nil {
				fmt.Println("TLS Version:", resp.TLS.Version)
				fmt.Println("TLS HandshakeComplete:",
					resp.TLS.HandshakeComplete)
				fmt.Println("TLS DidResume:", resp.TLS.DidResume)
				fmt.Println("TLS CipherSuite:", resp.TLS.CipherSuite)
				fmt.Println("TLS NegotiatedProtocol:",
					resp.TLS.NegotiatedProtocol)
				fmt.Println("TLS NegotiatedProtocolIsMutual:",
					resp.TLS.NegotiatedProtocolIsMutual)
				fmt.Println("TLS ServerName:", resp.TLS.ServerName)
			}
		}
	}
}
