package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	targetURL := os.Getenv("URL")
	if targetURL == "" {
		log.Fatalf("URL not set")
	}

	target, err := url.Parse(targetURL)
	if err != nil {
		log.Fatalf("Failed to parse target URL: %v", err)
	}

	log.Printf(`Starting with the following settings
TargetURL: %s
Basic_auth: %s
Username: %s
Password: %t`,
		targetURL,
		os.Getenv("BASIC_AUTH"),
		os.Getenv("USERNAME"),
		func() bool {
			return len(os.Getenv("PASSWORD")) > 0
		}(),
	)

	proxy := httputil.NewSingleHostReverseProxy(target)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Add Basic Authentication header if requested
		if os.Getenv("BASIC_AUTH") == "true" {
			r.Header.Add(
				"Authorization",
				"Basic "+base64.StdEncoding.EncodeToString(
					[]byte(os.Getenv("USERNAME")+":"+os.Getenv("PASSWORD")),
				),
			)
		}

		r.Host = target.Host
		proxy.ServeHTTP(w, r)
	})

	server := &http.Server{Addr: ":8099"}
	log.Println("Reverse proxy running on :8099 →", targetURL)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
