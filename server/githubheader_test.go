package server

import (
	"net/http"
	"testing"
)

func httpRequestEmpty() *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/data", nil)
	return req
}

func httpRequestWithContentTypeApplicationJson() *http.Request {
	req := httpRequestEmpty()
	req.Header.Add("Content-Type", "application/json")
	return req
}

func TestGithubHeader_Init(t *testing.T) {
	var gitH = &GithubHeader{}
	emptyReq := httpRequestEmpty()

	gitH.Init(emptyReq)
	req := GithubHeader_Export_httpRequest(gitH)
	if emptyReq != req {
		t.Errorf(`Expected "%p" got "%p"`, emptyReq, req)
	}
	testPrivateField(t, string(GithubHeader_Export_event(gitH)))
	testPrivateField(t, GithubHeader_Export_hubSignatureSecret(gitH))
	testPrivateField(t, GithubHeader_Export_hubSignatureSecret256(gitH))
}

func testPrivateField(t *testing.T, methodOutput string) {
	if methodOutput != "" {
		t.Errorf(`Expected empty string, got "%s"`, methodOutput)
	}
}

func TestGithubHeader_isContentTypeApplicationJson(t *testing.T) {
	var gitH = &GithubHeader{}
	emptyReq := httpRequestEmpty()
	jsonReq := httpRequestWithContentTypeApplicationJson()

	gitH.Init(emptyReq)
	if GithubHeader_Export_IsContentTypeApplicationJson(gitH) {
		t.Error(`Expected false got true`)
	}
	gitH.Init(jsonReq)
	if GithubHeader_Export_IsContentTypeApplicationJson(gitH) == false {
		t.Error(`Expected true got false`)
	}
}
