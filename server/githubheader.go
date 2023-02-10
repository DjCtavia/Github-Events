package server

import (
	"net/http"
)

type GithubHeader struct {
	httpRequest        *http.Request
	hubSignatureSecret string
}

func NewGithubHeader(req *http.Request) *GithubHeader {
	gitH := &GithubHeader{}
	gitH.httpRequest = req
	gitH.hubSignatureSecret = req.Header.Get("X-Hub-Signature-256")
	return gitH
}

func (gitH *GithubHeader) IsValidHeader() bool {
	if gitH.isNotContentTypeApplicationJson() && gitH.isNotContentTypeFormURLEnconded() {
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

func (gitH *GithubHeader) isContentTypeFormURLEnconded() bool {
	contentType := gitH.httpRequest.Header.Get("Content-Type")
	if contentType == "application/x-www-form-urlencoded" {
		return true
	}
	return false
}

func (gitH *GithubHeader) isNotContentTypeFormURLEnconded() bool {
	return !gitH.isContentTypeFormURLEnconded()
}

func (gitH *GithubHeader) GetSignatureHeader() string {
	return gitH.hubSignatureSecret
}
