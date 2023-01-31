package request

import "net/http"

type RequestBuilder interface {
	MakeRequest() *http.Request
}

type RequestDirector struct {
	RequestBuilder RequestBuilder
	request        *http.Request
}

func (requestDirector *RequestDirector) ChangeBuilder(newRequestBuilder RequestBuilder) {
	requestDirector.RequestBuilder = newRequestBuilder
}

func (requestDirector *RequestDirector) MakeRequest() *http.Request {
	if requestDirector.RequestBuilder == nil {
		return nil
	}
	requestDirector.request = requestDirector.RequestBuilder.MakeRequest()
	return requestDirector.request
}

func (requestDirector *RequestDirector) GetRequest() *http.Request {
	return requestDirector.request
}
