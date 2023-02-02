package discord

type DiscordEmbed struct {
	Title       string         `json:"title,omitempty"`
	Type        string         `json:"type,omitempty"`
	Description string         `json:"description,omitempty"`
	Url         string         `json:"url,omitempty"`
	Timestamp   string         `json:"timestamp,omitempty"`
	Color       int            `json:"color,omitempty"`
	Footer      NotImplemented `json:"footer,omitempty"`
	Image       NotImplemented `json:"image,omitempty"`
	Thumbnail   NotImplemented `json:"thumbnail,omitempty"`
	Video       NotImplemented `json:"video,omitempty"`
	Provider    NotImplemented `json:"provider,omitempty"`
	Author      NotImplemented `json:"author,omitempty"`
	Fields      NotImplemented `json:"fields,omitempty"`
}
