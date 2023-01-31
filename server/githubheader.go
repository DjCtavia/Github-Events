package server

import "net/http"

type GithubHeader struct {
	httpRequest           *http.Request
	event                 GithubEvent
	hubSignatureSecret    string
	hubSignatureSecret256 string
}

func (gitH *GithubHeader) Init(req *http.Request) {
	gitH.httpRequest = req
	gitH.hubSignatureSecret = req.Header.Get("X-Hub-Signature")
	gitH.hubSignatureSecret = req.Header.Get("X-Hub-Signature-256")
	gitH.event = GithubEvent(req.Header.Get("X-GitHub-Event"))
}

func (gitH *GithubHeader) IsValidHeader() bool {
	if gitH.isNotContentTypeApplicationJson() {
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

func (gitH *GithubHeader) GetEvent() GithubEvent {
	return gitH.event
}
