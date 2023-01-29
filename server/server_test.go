package server

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"testing"
)

const (
	HTTPLOCALHOST = "http://localhost"
	EXPECTEDGOT   = `Expected: "%s" | Got: "%s"`
	ERROR         = `Error: %v`
	GOODMETHOD    = `Good method`
	BADMETHOD     = `Bad method`
)

// Server must be run aside for testing otherwise will fail
func TestRunServer(t *testing.T) {
	testPostMethod(t)
}

func testPostMethod(t *testing.T) {
	client := &http.Client{}
	jsonBody := []byte(`{}`)
	bodyReader := bytes.NewReader(jsonBody)
	url := fmt.Sprintf("%s:%d/data", HTTPLOCALHOST, DefaultServerPort)
	req, _ := http.NewRequest(http.MethodPost, url, bodyReader)

	resp, errResp := client.Do(req)
	if errResp != nil {
		t.Fatalf(ERROR, errResp)
	}
	defer req.Body.Close()
	bodyByte, err := io.ReadAll(resp.Body)

	if errResp != nil {
		t.Fatalf(ERROR, err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf(EXPECTEDGOT, getStatusString(http.StatusOK), getStatusString(resp.StatusCode))
	}
	if string(bodyByte) != GOODMETHOD {
		t.Fatalf(EXPECTEDGOT, GOODMETHOD, bodyByte)
	}
	t.Logf(EXPECTEDGOT, GOODMETHOD, bodyByte)
}

func getStatusString(statusCode int) string {
	return fmt.Sprintf("status %d", statusCode)
}
