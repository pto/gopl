// Ex09 outputs the content found at a URL and the HTTP status code.
package main

import (
	"crypto/tls"
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
				fmt.Println("TLS Version:",
					tlsVersion(resp.TLS.Version))
				fmt.Println("TLS HandshakeComplete:",
					resp.TLS.HandshakeComplete)
				fmt.Println("TLS DidResume:",
					resp.TLS.DidResume)
				fmt.Println("TLS CipherSuite:",
					cipherSuite(resp.TLS.CipherSuite))
				fmt.Println("TLS NegotiatedProtocol:",
					resp.TLS.NegotiatedProtocol)
				fmt.Println("TLS NegotiatedProtocolIsMutual:",
					resp.TLS.NegotiatedProtocolIsMutual)
			}
		}
	}
}

func cipherSuite(cs uint16) string {
	switch cs {
	case tls.TLS_RSA_WITH_RC4_128_SHA:
		return "TLS_RSA_WITH_RC4_128_SHA"
	case tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA:
		return "TLS_RSA_WITH_RC4_128_SHA"
	case tls.TLS_RSA_WITH_AES_128_CBC_SHA:
		return "TLS_RSA_WITH_AES_128_CBC_SHA"
	case tls.TLS_RSA_WITH_AES_256_CBC_SHA:
		return "TLS_RSA_WITH_AES_256_CBC_SHA"
	case tls.TLS_RSA_WITH_AES_128_GCM_SHA256:
		return "TLS_RSA_WITH_AES_128_GCM_SHA256"
	case tls.TLS_RSA_WITH_AES_256_GCM_SHA384:
		return "TLS_RSA_WITH_AES_256_GCM_SHA384"
	case tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA:
		return "TLS_ECDHE_ECDSA_WITH_RC4_128_SHA"
	case tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA:
		return "TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA"
	case tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA:
		return "TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA"
	case tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA:
		return "TLS_ECDHE_RSA_WITH_RC4_128_SHA"
	case tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA:
		return "TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA"
	case tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA:
		return "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA"
	case tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA:
		return "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA"
	case tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:
		return "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"
	case tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256:
		return "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256"
	case tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384:
		return "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"
	case tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384:
		return "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"
	case tls.TLS_FALLBACK_SCSV:
		return "TLS_FALLBACK_SCSV"
	default:
		return "(unknown)"
	}
}

func tlsVersion(tv uint16) string {
	switch tv {
	case tls.VersionSSL30:
		return "VersionSSL30"
	case tls.VersionTLS10:
		return "VersionTLS10"
	case tls.VersionTLS11:
		return "VersionTLS11"
	case tls.VersionTLS12:
		return "VersionTLS12"
	default:
		return "(unknown)"
	}
}
