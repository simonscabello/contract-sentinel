package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthHandler)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("esperado status 200, obteve %d", rr.Code)
	}

	if rr.Body.String() != "OK\n" {
		t.Errorf("esperado corpo 'OK', obteve %q", rr.Body.String())
	}
}
