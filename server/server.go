package server

import (
	"fmt"
	"log"
	"net/http"
)

const DefaultServerPort = 8080

type Server struct {
	Port int
}

func (serv Server) SetPort(port int) {
	serv.Port = port
}

func (serv Server) RunServer() {
	fmt.Printf("Server is Running on Port %d.\n", serv.Port)
	http.HandleFunc("/data", DataHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", serv.Port), nil))
}

func DataHandler(res http.ResponseWriter, req *http.Request) {
	if isNotRequestPostMethod(req) {
		fmt.Fprintf(res, "Bad method")
		return
	}
	if req == nil {
		fmt.Fprintf(res, "Req is empty")
	}
	headerHandler := GithubHeader{}
	headerHandler.Init(req)

	if headerHandler.IsValidHeader() == false {
		fmt.Fprintf(res, "Bad headers")
		return
	}

	_, err := res.Write([]byte{})
	if err != nil {
		log.Fatalf("Error when writing bytes: %v", err)
	}
	fmt.Fprintf(res, "Good method")
}

func IsRequestPostMethod(req *http.Request) bool {
	return req.Method == http.MethodPost
}

func isNotRequestPostMethod(req *http.Request) bool {
	return !IsRequestPostMethod(req)
}
