package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type headerTest struct {
	Key   string
	Value string
}

type dataHandlerTestsStruct struct {
	name               string
	StatusCodeExpected int
	Method             string
	Headers            []headerTest
	Body               io.Reader
}

var dataHandlerTests = []dataHandlerTestsStruct{
	{
		"Using bad method",
		http.StatusMethodNotAllowed,
		http.MethodGet,
		nil,
		strings.NewReader("{}"),
	},
	{
		"Using bad headers",
		http.StatusBadRequest,
		http.MethodPost,
		[]headerTest{
			{"Content-Type", "text/plain"},
		},
		strings.NewReader("{}"),
	},
	{
		"Using bad signature",
		http.StatusUnauthorized,
		http.MethodPost,
		[]headerTest{
			{"Content-Type", "application/json"},
			{"X-Hub-Signature-256", "fakehash"},
		},
		strings.NewReader("{}"),
	},
	{
		"Well formated request",
		http.StatusOK,
		http.MethodPost,
		[]headerTest{
			{"Content-Type", "application/json"},
			{"X-Hub-Signature-256", "sha256=7b7d8759ecadd90d4ff0f0fa1ae3f59105c13b4a56e7ca3ae970580a56a70ca39dc7"},
		},
		strings.NewReader("{}"),
	},
}

// Server must be run aside for testing otherwise will fail
func TestDataHandler(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(DataHandler))
	client := http.DefaultClient

	for _, test := range dataHandlerTests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(test.Method, server.URL, test.Body)
			if err != nil {
				t.Error(err)
			}

			for _, header := range test.Headers {
				req.Header.Set(header.Key, header.Value)
			}

			res, err := client.Do(req)
			if err != nil {
				t.Error(err)
			}
			if res.StatusCode != test.StatusCodeExpected {
				t.Errorf("Expected status code: %d, Got: %d\n", test.StatusCodeExpected, res.StatusCode)
			}
		})
	}
}
