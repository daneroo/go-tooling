package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"daniel@golang.org", "gopher daniel"},
		{"someone-else", "dear"},
	}
	for _, c := range cases {
		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/"+c.in,
			nil,
		)
		if err != nil {
			t.Fatalf("Could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		handler(rec, req)
		if rec.Code != 200 {
			t.Errorf("Expected status 200; got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), c.out) {
			t.Errorf("unexpected body in response: %q", rec.Body.String())
		}

	}
}

func BenchmarkHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {

		req, err := http.NewRequest(
			http.MethodGet,
			"http://localhost:8080/daniel@golang.org",
			nil,
		)
		if err != nil {
			b.Fatalf("Could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		handler(rec, req)
		if rec.Code != 200 {
			b.Errorf("Expected status 200; got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "gopher daniel") {
			b.Errorf("unexpected body in response: %q", rec.Body.String())
		}

	}
}
