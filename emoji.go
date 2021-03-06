package slack

import (
	"net/url"
)

type emojiResponseFull struct {
	Emoji map[string]string `json:"emoji"`
	SlackResponse
}

// GetEmoji retrieves all the emojis
func (api *Slack) GetEmoji() (map[string]string, error) {
	values := url.Values{
		"token": {api.config.token},
	}
	response := &emojiResponseFull{}
	err := parseResponse("emoji.list", values, response, api.debug)
	if err != nil {
		return nil, err
	}
	if !response.Ok {
		return nil, NewSlackError(response.Error)
	}
	return response.Emoji, nil
}
