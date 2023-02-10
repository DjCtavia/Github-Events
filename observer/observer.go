package observer

import (
	"Discord"
	"djctavia/server"
	"fmt"
)

type Observer interface {
	update(event *server.GithubEvent)
	GetID() string
}

type GithubObserver struct {
	Id   string
	Hook *discord.DiscordHook
}

func (gho *GithubObserver) update(event *server.GithubEvent) {
	for _, commit := range event.Commits {
		fmtMessage := fmt.Sprintf("A new commit with ID: `%s` has been pushed!", commit.ID)
		dThumbnail := discord.DiscordThumbnail{Url: event.Sender.AvatarURL, Height: 16, Width: 16}
		dEmbed := discord.DiscordEmbed{
			Type:        "rich",
			Title:       fmt.Sprintf("Commit - %s", commit.ID),
			Url:         commit.URL,
			Description: fmtMessage,
			Thumbnail:   dThumbnail,
			Timestamp:   commit.Timestamp,
		}
		dMessage := discord.DiscordMessage{Embeds: []discord.DiscordEmbed{dEmbed}}

		gho.Hook.SendMessage(dMessage)
	}
}

func (gho *GithubObserver) GetID() string {
	return gho.Id
}
