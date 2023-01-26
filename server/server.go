package server

import (
	"fmt"
	"log"
	"net/http"
)

const ServerPort = 8080

func RunServer() {
	fmt.Printf("Server is Running on port %d.\n", ServerPort)
	http.HandleFunc("/data", DataHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", ServerPort), nil))
}

func DataHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello Github! Welcome on \"%s\" path!", req.URL.Path[1:])
}
