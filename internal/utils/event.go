package utils

import (
	"encoding/json"
	"errors"
	"io/fs"
)

// GitHub issue
type issue struct {
	Number int `json:"number"`
}

// event that triggers the action
type event struct {
	// action name
	Action string `json:"action"`
	// pull request number
	Number int `json:"number"`
	// GitHub issue, if the event is issue
	Issue issue `json:"issue"`
}

// ReadEvent reads and parse event meta definition
func GetEventNumber(fsys fs.FS) (int, error) {
	file, err := fsys.Open(
		ReadEnvString("GITHUB_EVENT_PATH")[1:],
	)

	if err != nil {
		return 0, errors.New("Failed to read event file")
	}

	var event event

	if err := json.NewDecoder(file).Decode(&event); err != nil {
		return 0, errors.New("Failed to parse event file")
	}

	if event.Number > 0 {
		return event.Number, nil
	}

	return event.Issue.Number, nil
}
