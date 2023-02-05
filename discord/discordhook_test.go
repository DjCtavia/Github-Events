package discord

import "testing"

func TestDiscordHook_SendMessage(t *testing.T) {
	dh := DiscordHook{"https://discord.com/api/webhooks/1070751564505038990/3FbPB1qPAWKf_ISUsxyg3ZE-hF-_oyfdGArrdfTGQXuffAvpjWRoE4k69yPsE_mXkiOi"}

	testDiscordHook_SendMessage_content(t, &dh)
	testDiscordHook_SendMessage_embed(t, &dh)
}

func testDiscordHook_SendMessage_content(t *testing.T, dh *DiscordHook) {
	dmessage := DiscordMessage{}
	err := dh.SendMessage(dmessage)
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	dmessage = DiscordMessage{Content: "It shoud work!"}
	err = dh.SendMessage(dmessage)
	if err != nil {
		t.Error(err)
		return
	}
}

func testDiscordHook_SendMessage_embed(t *testing.T, dh *DiscordHook) {
	dthumbnail := DiscordThumbnail{Url: "https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png", Width: 16, Height: 16}
	dembed := DiscordEmbed{
		Type:        "rich",
		Title:       "Test",
		Description: "Description test",
		Color:       HexRGBToInt("6fab54"),
		Thumbnail:   dthumbnail,
	}
	embeds := []DiscordEmbed{dembed}
	dmessage := DiscordMessage{Embeds: embeds}

	err := dh.SendMessage(dmessage)
	if err != nil {
		t.Error(err)
		return
	}
}
