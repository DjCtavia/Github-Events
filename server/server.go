package server

import (
	"fmt"
	"log"
	"net/http"
)

func RunServer() {
	fmt.Println("Server is Running on port 8080.")
	http.HandleFunc("/data", DataHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func DataHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello Github! Welcome on \"%s\" path!", req.URL.Path[1:])
}
