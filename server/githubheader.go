package server

import (
	"net/http"
)

type GithubHeader struct {
	SecretToken        string
	httpRequest        *http.Request
	hubSignatureSecret string
}

func (gitH *GithubHeader) Init(req *http.Request) {
	gitH.httpRequest = req
	gitH.hubSignatureSecret = req.Header.Get("X-Hub-Signature")
}

func (gitH *GithubHeader) IsValidHeader() bool {
	if gitH.isNotContentTypeApplicationJson() {
		return false
	}
	if gitH.hubSignatureSecret == "" {
		return false
	}
	return true
}

func (gitH *GithubHeader) isContentTypeApplicationJson() bool {
	contentType := gitH.httpRequest.Header.Get("Content-Type")
	if contentType == "application/json" {
		return true
	}
	return false
}

func (gitH *GithubHeader) isNotContentTypeApplicationJson() bool {
	return !gitH.isContentTypeApplicationJson()
}

func (gitH *GithubHeader) GetSignatureHeader() string {
	return gitH.hubSignatureSecret
}
