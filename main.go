package main

import (
	"djctavia/server"
	"fmt"
)

func main() {
	fmt.Println("GithubEvents Started.")
	serv := server.Server{Port: server.DEFAULTSERVERPORT}
	serv.RunServer()
}
