package request

import (
	"log"
	"net/http"
)

type RequestEmpty struct{}

func (requestEmpty *RequestEmpty) MakeRequest() *http.Request {
	req, err := http.NewRequest(http.MethodPost, "/data", nil)
	if err != nil {
		log.Fatal("Couldn't create request")
	}
	return req
}
