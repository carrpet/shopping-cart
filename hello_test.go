package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("homeHandler returned wrong status code: got %v, want %v",
			status, http.StatusOK)
	}

	expectedCT := "text/html"
	ct := rr.Header().Get("Content-Type")
	if ct != expectedCT {
		t.Errorf("handler return unexpected header content-type: got %v want %v", ct, expectedCT)
	}
}
