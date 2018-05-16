package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCheeseburgerOK(t *testing.T) {
	// The handler is expecting a Request and a ResponseWriter
	// as parameters

	// Create a requst that matches the handler we're
	req, err := http.NewRequest("GET", "/ichc", nil)
	if err != nil {
		t.Fatal(err)
	}
	// ResponseRecoder apparently satisfies http.ResponseWriter:
	// https://blog.questionable.services/article/testing-http-handlers-go/
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCheeseburger)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
func TestGetCheeseburgerResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/ichc", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCheeseburger)

	handler.ServeHTTP(rr, req)

	expected := "CheeseburgerCheeseburger"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
