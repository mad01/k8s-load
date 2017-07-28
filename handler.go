package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const (
	minIntVal = 50000
	maxIntVal = 1000000
)

func randomRange(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max-min) + min
}

// Fib get fib for n
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func httpFibHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	num := randomRange(minIntVal, maxIntVal)
	for n := 0; n < num; n++ {
		Fib(10)
	}
	elapsed := time.Since(start)
	fmt.Fprintf(w, "fib(10) loops n: %v elapsed: %s", num, elapsed)
}

func httpHostnameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server hostname: %s", hostname)
}
