package request

import "net/http"

type RequestBuilder interface {
	MakeRequest() *http.Request
}

type RequestDirector struct {
	requestBuilder *RequestBuilder
	request        *http.Request
}

func (requestDirector *RequestDirector) ChangeBuilder(newRequestBuilder *RequestBuilder) {
	requestDirector.requestBuilder = newRequestBuilder
}

func (requestDirector *RequestDirector) MakeRequest() *http.Request {
	if requestDirector.request == nil {
		return nil
	}
	requestDirector.request = requestDirector.MakeRequest()
	return requestDirector.request
}

func (requestDirector *RequestDirector) GetRequest() *http.Request {
	return requestDirector.request
}
