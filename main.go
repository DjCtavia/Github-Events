package main

import (
	discord "Discord"
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

	registerHooks(&hooksFlag)

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

func registerHooks(hooksFlag *hookURLFlag) {
	ArrayOfDataHook, err := hooksFlag.GetArrayOfDataHook()
	if err != nil {
		os.Exit(64)
	}
	if len(ArrayOfDataHook) == 0 {
		fmt.Fprintf(os.Stderr, "Add more hooks with command -hookUrl, type -help for more details")
		os.Exit(64)
	}
	for _, hook := range ArrayOfDataHook {
		observer.GlobalSubject.Register(&observer.GithubObserver{Id: hook.ID, Hook: &discord.DiscordHook{Url: hook.hookURL}})
	}
}
