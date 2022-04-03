package repository

import (
	"encoding/json"
	"io/fs"
)

var (
	filenames = []string{
		"abbreviates.json",
		"brands.json",
		"general.json",
		"products.json",
		"softwares.json",
	}
)

// GetDictionary returns word-to-word mapping dictionary
func GetDictionary(fsys fs.FS) map[string]string {
	dict := make(map[string]string)

	for _, filename := range filenames {
		fileDict := parseDictionaryFile(fsys, filename)

		for k, v := range fileDict {
			dict[k] = v
		}
	}

	return dict
}

func parseDictionaryFile(fsys fs.FS, filename string) map[string]string {
	file, err := fsys.Open(filename)

	if err != nil {
		return map[string]string{}
	}

	var data map[string]string
	err = json.NewDecoder(file).Decode(&data)

	if err != nil {
		return map[string]string{}
	}

	return data
}
