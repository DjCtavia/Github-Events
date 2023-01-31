package server

import "net/http"

var GithubHeader_Export_httpRequest = func(header *GithubHeader) *http.Request {
	return header.httpRequest
}

var GithubHeader_Export_event = func(header *GithubHeader) GithubEvent {
	return header.event
}

var GithubHeader_Export_hubSignatureSecret = func(header *GithubHeader) string {
	return header.hubSignatureSecret
}

var GithubHeader_Export_hubSignatureSecret256 = func(header *GithubHeader) string {
	return header.hubSignatureSecret256
}

// methods
var GithubHeader_Export_IsContentTypeApplicationJson = (*GithubHeader).isContentTypeApplicationJson
var GithubHeader_Export_IsNotContentTypeApplicationJson = (*GithubHeader).isContentTypeApplicationJson
