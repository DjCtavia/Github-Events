package server

import "net/http"

var GithubHeader_Export_httpRequest = func(header *GithubHeader) *http.Request {
	return header.httpRequest
}

var GithubHeader_Export_hubSignatureSecret = func(header *GithubHeader) string {
	return header.hubSignatureSecret
}

// methods
var GithubHeader_Export_IsContentTypeApplicationJson = (*GithubHeader).isContentTypeApplicationJson
var GithubHeader_Export_IsNotContentTypeApplicationJson = (*GithubHeader).isContentTypeApplicationJson
