package request

import (
	"bytes"
	"log"
	"net/http"
)

type RequestEmptyJson struct{}

func (requestEmptyJson *RequestEmptyJson) MakeRequest() *http.Request {
	jsonData := []byte("{}")
	req, err := http.NewRequest(http.MethodPost, "/data", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("Couldn't create request")
	}
	req.Header.Add("Content-Type", "application/json")
	return req
}
