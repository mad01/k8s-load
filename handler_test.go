package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/fib", nil)
	if err != nil {
		t.Errorf("Failed to create new request: %v", err)
	}
	rec := httptest.NewRecorder()
	httpFibHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code: %v got %v", http.StatusOK, rec.Code)
	}
}
