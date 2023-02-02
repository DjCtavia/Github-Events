package discord

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DiscordError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (dErr DiscordError) Error() string {
	return fmt.Sprintf("%s | Code: %d", dErr.Message, dErr.Code)
}

func CreateDiscordError(r *http.Response) DiscordError {
	bodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		return DiscordError{"Can't read body", -1}
	}
	dErr := DiscordError{}
	err = json.Unmarshal(bodyByte, &dErr)
	if err != nil {
		return DiscordError{"Can't Unmarshal Body as JSON", -2}
	}
	return dErr
}
