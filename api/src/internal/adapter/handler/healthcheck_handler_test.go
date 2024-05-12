package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheckHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/healtcheck", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("ステータスコードが異なります： got: %v want %v", status, http.StatusOK)
	}

	expected := "ok"
	if rr.Body.String() != expected {
		t.Errorf("返却値が異なります： got: %v want %v", rr.Body.String(), expected)
	}
}
