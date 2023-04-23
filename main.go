package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Create HTTP server with autocert TLS configuration

	cert, err := tls.LoadX509KeyPair("TLS CERT HERE", "TLS CERT KEY HERE")
	if err != nil {
		fmt.Println("Failed to load SSL certificate and key pair:", err)
		return
	}

	// Create TLS configuration with the loaded certificate and key pair
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	// Create HTTPS server with custom TLS configuration
	s := &http.Server{
		Addr:      ":443",
		Handler:   http.NewServeMux(),
		TLSConfig: tlsConfig,
	}

	// Create a new ServeMux for handling requests
	mux := http.NewServeMux()

	mux.HandleFunc("/.well-known/webfinger", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{
                        "subject" : "acct:youremail@example.com",
                        "links" :
                        [
                                {
                                        "rel" : "http://openid.net/specs/connect/1.0/issuer",
                                        "href" : "OCID ISSUER URL HERE"
                                }
                        ]
                }`))
	})

	// Register the ServeMux as the handler for the HTTPS server
	s.Handler = mux

	// Configure logger to output to console
	log.SetOutput(os.Stdout)

	// Start the HTTP server with TLS
	err = s.ListenAndServeTLS("", "")
	if err != nil {
		fmt.Println("Failed to start HTTPS server:", err)
	}
}
