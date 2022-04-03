package repository

import (
	"encoding/json"
	"fmt"
	"io/fs"
)

// GetDictionary returns word-to-word mapping dictionary
func GetDictionary(fsys fs.FS, files []string) map[string]string {
	dict := make(map[string]string)

	for _, filename := range files {
		fileDict := parseDictionaryFile(fsys, filename)

		for k, v := range fileDict {
			dict[k] = v
		}
	}

	return dict
}

func parseDictionaryFile(fsys fs.FS, filename string) map[string]string {
	file, err := fsys.Open(filename)

	fmt.Println(err)

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
