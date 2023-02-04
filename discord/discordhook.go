package discord

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type DiscordHook struct {
	Url string
}

// https://discord.com/developers/docs/resources/channel#create-message
func (dh *DiscordHook) SendMessage(message DiscordMessage) error {
	messageJSON, err := dh.convertMessageToJSON(message)
	if err != nil {
		return err
	}
	req, _ := dh.createRequest(messageJSON)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if dh.isNotValidStatusCode(res) {
		return CreateDiscordError(res)
	}
	return nil
}

func (dh *DiscordHook) convertMessageToJSON(message DiscordMessage) ([]byte, error) {
	messageJSON, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}
	return messageJSON, nil
}

func (dh *DiscordHook) createRequest(jsonByteData []byte) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, dh.Url, bytes.NewBuffer(jsonByteData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (dh *DiscordHook) isValidStatusCode(res *http.Response) bool {
	return res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent
}

func (dh *DiscordHook) isNotValidStatusCode(res *http.Response) bool {
	return dh.isValidStatusCode(res) == false
}
