package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"log"
	"net/http"
)

func main() {
	limiter := rate.NewLimiter(10, 1)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		fmt.Fprintln(w, "Hello, World")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
