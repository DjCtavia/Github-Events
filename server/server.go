package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

const DEFAULTSERVERPORT = 8080

type Server struct {
	Port int
}

func NewServer(port int) *Server {
	return &Server{Port: port}
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
	defer req.Body.Close()
	if req.Method != http.MethodPost {
		http.Error(res, "Bad method", http.StatusMethodNotAllowed)
		return
	}

	headerHandler := NewGithubHeader(req)
	if headerHandler.IsValidHeader() == false {
		http.Error(res, "Bad headers", http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Server error", http.StatusInternalServerError)
		return
	}

	gitSignature := GithubSignature{DEFAULTSIGNATURE}
	if gitSignature.Verify(body, headerHandler.GetSignatureHeader()) != nil {
		http.Error(res, "Signature not matching", http.StatusUnauthorized)
		return
	}

	var event GitHubEvent
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&event)
	if err != nil {
		http.Error(res, "Server error", http.StatusInternalServerError)
		return
	}
}
