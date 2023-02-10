package main

import (
	"djctavia/server"
	"fmt"
	"observer"
)

func main() {
	fmt.Println("GithubEvents Started.")

	// Example
	// observer.GlobalSubject.Register(&observer.GithubObserver{Id: "Main Discord Server", Hook: &discord.DiscordHook{Url: "https://discord.com/api/webhooks/7777777777777777777/abcdefg_abcd_abcdefg012345-abcedfg124567987a_bcedf_abcedf01234456789"}})

	serv := server.GetServer()
	serv.BindPort(server.DEFAULTSERVERPORT)
	serv.BindNotifyFunc(NotifyEvent)
	serv.RunServer()
}

func NotifyEvent(event *server.GithubEvent) {
	go observer.GlobalSubject.Notify(event)
}
