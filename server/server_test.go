package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Server must be run aside for testing otherwise will fail
func TestDataHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(DataHandler))
	resp, err := http.Post(server.URL, "application/json", strings.NewReader("{}"))
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected %d status, got %d", http.StatusOK, resp.StatusCode)
	}
	defer resp.Body.Close()

	expectedBody := "Good method"
	byteBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	if string(byteBody) != expectedBody {
		t.Errorf(`Expected "%s" body, got "%s`, expectedBody, string(byteBody))
	}
}
