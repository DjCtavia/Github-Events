package discord

type NotImplemented string

// https://discord.com/developers/docs/resources/channel#embed-object
type DiscordMessage struct {
	Content          string         `json:"content,omitempty"`
	Nonce            int            `json:"nonce,omitempty"`
	Tts              bool           `json:"tts,omitempty"`
	Embeds           []DiscordEmbed `json:"embeds,omitempty"`
	AllowedMentions  string         `json:"allowed_mentions,omitempty"`
	MessageReference NotImplemented `json:"message_reference,omitempty"`
	Components       NotImplemented `json:"components,omitempty"`
	StickerIds       NotImplemented `json:"sticker_ids,omitempty"`
	Files            NotImplemented `json:"files[n],omitempty"`
	PayloadJSON      string         `json:"payload_json,omitempty"`
	Attachments      NotImplemented `json:"attachments,omitempty"`
	Flags            int            `json:"flags,omitempty"`
}
