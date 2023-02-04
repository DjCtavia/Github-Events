package discord

import "testing"

func TestDiscordHook_SendMessage(t *testing.T) {
	dh := DiscordHook{"https://discord.com/api/webhooks/1070751564505038990/3FbPB1qPAWKf_ISUsxyg3ZE-hF-_oyfdGArrdfTGQXuffAvpjWRoE4k69yPsE_mXkiOi"}

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
