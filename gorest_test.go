package main

import (
	"strings"
	"net/http"
	"net/http/httptest"
	"net/url"
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
func TestCreateCheeseburgerOK(t *testing.T) {
	req, err := http.NewRequest("POST", "/ichc", nil)

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCheeseburger)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// THIS TEST IS NOT WORKING
// I do not think it is passing query parameters,
// or I have been unable to capture them
// I cannot confirm that the custom header is being
// set, either
func TestCreateCheeseburgerResponse(t *testing.T) {
	query_params := url.Values{}
	query_params.Set("toppings", "lettuce")
	query_params.Add("cheese", "true")

	req, err := http.NewRequest("POST", "/ichc",
		strings.NewReader(query_params.Encode()))
	req.Header.Add("X-gorest_test", "True")

	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCheeseburger)

	handler.ServeHTTP(rr, req)

	expected := "yo"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestGetIndexOK(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetIndex)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
func TestGetIndexResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/ichc", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetIndex)

	handler.ServeHTTP(rr, req)

  expected := "There's not really supposed to be an index"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
