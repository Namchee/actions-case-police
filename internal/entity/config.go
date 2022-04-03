package entity

import (
	"encoding/json"
	"errors"

	"github.com/Namchee/actions-case-police/internal/utils"
)

var (
	defaultPreset = []string{
		"abbreviates",
		"brands",
		"general",
		"products",
		"softwares",
	}
)

type Configuration struct {
	Token      string
	Fix        bool
	Preset     []string
	Dictionary map[string]string
}

// ReadConfiguration read and parse action user input
// from workflow definition
func ReadConfiguration() (*Configuration, error) {
	token := utils.ReadEnvString("INPUT_GITHUB_TOKEN")

	if token == "" {
		return nil, errors.New("Missing GitHub token")
	}

	config := &Configuration{
		Token:  token,
		Fix:    utils.ReadEnvBool("INPUT_FIX"),
		Preset: defaultPreset,
	}

	preset := utils.ReadEnvStringArray("INPUT_PRESET")
	if len(preset) > 0 {
		config.Preset = preset
	}

	customDict := utils.ReadEnvString("INPUT_DICTIONARY")
	var dictionary map[string]string

	err := json.Unmarshal([]byte(customDict), &dictionary)
	if err == nil {
		config.Dictionary = dictionary
	}

	return config, nil
}
