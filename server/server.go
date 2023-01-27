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
	fmt.Fprintf(res, "Hello Github! Welcome on \"%s\" path!", req.URL.Path[1:])
	if !IsRequestPostMethod(req) {
		return
	}
	fmt.Fprintf(res, "\nYou're using the POST method!")

}

func IsRequestPostMethod(req *http.Request) bool {
	return req.Method == http.MethodPost
}
