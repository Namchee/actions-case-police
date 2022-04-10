package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"

	"github.com/Namchee/actions-case-police/internal/utils"
)

// Event that triggers the action
type Event struct {
	// action name
	Action string `json:"action"`
	// pull request number
	Number int `json:"number"`
}

// ReadEvent reads and parse event meta definition
func ReadEvent(fsys fs.FS) (*Event, error) {
	file, err := fsys.Open(
		utils.ReadEnvString("GITHUB_EVENT_PATH")[1:],
	)

	if err != nil {
		return nil, errors.New("Failed to read event file")
	}

	var event Event

	b, err := ioutil.ReadAll(file)
	fmt.Println(b)

	if err := json.NewDecoder(file).Decode(&event); err != nil {
		return nil, errors.New("Failed to parse event file")
	}

	return &event, nil
}
