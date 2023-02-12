package main

import (
	"djctavia/server"
	"flag"
	"fmt"
	"observer"
	"os"
)

func main() {
	fmt.Println("GithubEvents Started.")

	var hooksFlag hookURLFlag
	flag.Var(&hooksFlag, "hookUrl", "add a discord hook to attach: uniqueID@hookURL")
	helpFlag := flag.Bool("help", false, "show help")
	runningPortFlag := flag.Int("server port", server.DEFAULTSERVERPORT, "server port for receiving github events")

	flag.Parse()
	isHelpFlagPresent(helpFlag)

	// Example
	// observer.GlobalSubject.Register(&observer.GithubObserver{Id: "Main Discord Server", Hook: &discord.DiscordHook{Url: "https://discord.com/api/webhooks/7777777777777777777/abcdefg_abcd_abcdefg012345-abcedfg124567987a_bcedf_abcedf01234456789"}})

	serv := server.GetServer()
	serv.BindPort(*runningPortFlag)
	serv.BindNotifyFunc(NotifyEvent)
	serv.RunServer()
}

func NotifyEvent(event *server.GithubEvent) {
	go observer.GlobalSubject.Notify(event)
}

func isHelpFlagPresent(helpFlagEnabled *bool) {
	if *helpFlagEnabled {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}
}
