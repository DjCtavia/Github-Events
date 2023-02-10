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
			{"X-Hub-Signature-256", "sha256=9fb90896e56bbfe95fa925eb1ef4e7228d7a3d415f54a2dde5aa35ee118e2b8f"},
		},
		strings.NewReader("{}"),
	},
}

// JSON is Minify
const COMMITOBJECT string = `{"action":"commit","repository":{"name":"Github-Events","owner":{"login":"DjCtavia"}},"sender":{"login":"DjCtavia"},"commit":{"sha":"9ece3c053eb8b649a74a7fced0f51aa69ae66763","message":"Server module init","url":"https://github.com/DjCtavia/Github-Events/commit/9ece3c053eb8b649a74a7fced0f51aa69ae66763"}}`

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

	t.Run("CommitPayload", testDataHandler_WithCommitPayload)
}

func testDataHandler_WithCommitPayload(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(DataHandler))
	client := http.DefaultClient
	req, err := http.NewRequest(http.MethodPost, server.URL, strings.NewReader(COMMITOBJECT))
	if err != nil {
		t.Error(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hub-Signature-256", "sha256=32b658bc7e95705cedb3c5ac88d8beda18ac6c4705c1e28d96d16105bc6c42d3")
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status code: %d, Got: %d\n", http.StatusOK, res.StatusCode)
	}
}
