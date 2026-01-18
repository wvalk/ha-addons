package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, err := url.Parse("http://192.168.1.151")
	if err != nil {
		log.Fatalf("Failed to parse target URL: %v", err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	// Optional: Modify request before forwarding
	proxy.ModifyResponse = func(resp *http.Response) error {
		// Add custom response handling if needed
		return nil
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.Host = target.Host
		proxy.ServeHTTP(w, r)
	})

	server := &http.Server{
		Addr: ":8080",
	}

	log.Println("Reverse proxy running on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
