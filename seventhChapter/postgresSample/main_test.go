package main_test

import (
	"net/http"
	"testing"
)

func TestGetOrginalLink(t *testing.T) {
	response, err := http.Get("http://localhost:8000/v1/short/1/")
	if response.StatusCode != http.StatusOK {
		t.Errorf("Expected response code %d. Got %d\n", http.StatusOK, response.StatusCode)
	}
	if err != nil {
		t.Errorf("error happend: %s", err.Error())
	}
}
