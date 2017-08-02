package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkHandlerFib(b *testing.B) {
	for i := 0; i < b.N; i++ {

		req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/fib", nil)
		if err != nil {
			b.Errorf("Failed to create new request: %v", err)
		}
		rec := httptest.NewRecorder()
		httpFibHandler(rec, req)

		if rec.Code != http.StatusOK {
			b.Errorf("Expected status code: %v got %v", http.StatusOK, rec.Code)
		}
	}
}
