package server

import (
	"testing"
	"tests/request"
)

func TestGithubHeader_Init(t *testing.T) {
	var gitH *GithubHeader
	reqDir := request.RequestDirector{RequestBuilder: &request.RequestEmpty{}}
	emptyReq := reqDir.MakeRequest()

	gitH = NewGithubHeader(emptyReq)
	req := GithubHeader_Export_httpRequest(gitH)
	if emptyReq != req {
		t.Errorf(`Expected "%p" got "%p"`, emptyReq, req)
	}
	testPrivateField(t, GithubHeader_Export_hubSignatureSecret(gitH))
}

func testPrivateField(t *testing.T, methodOutput string) {
	if methodOutput != "" {
		t.Errorf(`Expected empty string, got "%s"`, methodOutput)
	}
}

func TestGithubHeader_isContentTypeApplicationJson(t *testing.T) {
	var gitH *GithubHeader
	reqDir := request.RequestDirector{RequestBuilder: &request.RequestEmpty{}}
	emptyReq := reqDir.MakeRequest()
	reqDir.ChangeBuilder(&request.RequestEmptyJson{})
	jsonReq := reqDir.MakeRequest()

	gitH = NewGithubHeader(emptyReq)
	if GithubHeader_Export_IsContentTypeApplicationJson(gitH) {
		t.Error(`Expected false got true`)
	}
	gitH = NewGithubHeader(jsonReq)
	if GithubHeader_Export_IsContentTypeApplicationJson(gitH) == false {
		t.Error(`Expected true got false`)
	}
}
