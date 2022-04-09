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
	Exclude    []string
	Dictionary map[string]string
}

// ReadConfiguration read and parse action user input
// from workflow definition
func ReadConfiguration() (*Configuration, error) {
	token := utils.ReadEnvString("INPUT_ACCESS_TOKEN")
	if token == "" {
		return nil, errors.New("Missing GitHub token")
	}

	config := &Configuration{
		Token:  token,
		Fix:    utils.ReadEnvBool("INPUT_FIX"),
		Preset: defaultPreset,
	}

	preset := validatePresets(utils.ReadEnvStringArray("INPUT_PRESET"))
	if len(preset) > 0 {
		config.Preset = preset
	}

	customDict := utils.ReadEnvString("INPUT_DICTIONARY")
	var dictionary map[string]string

	exclusion := utils.ReadEnvStringArray("INPUT_EXCLUDE")
	if len(exclusion) > 0 {
		config.Exclude = exclusion
	}

	err := json.Unmarshal([]byte(customDict), &dictionary)
	if err == nil {
		config.Dictionary = dictionary
	}

	return config, nil
}

// validatePresets validates and return valid presets from user inputs
func validatePresets(presets []string) []string {
	var newPresets []string

	for _, preset := range presets {
		if utils.ContainsString(defaultPreset, preset) {
			newPresets = append(newPresets, preset)
		}
	}

	return newPresets
}
