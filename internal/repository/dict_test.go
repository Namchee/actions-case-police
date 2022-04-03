package repository

import (
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

func TestGetDictionary(t *testing.T) {
	type fileMock struct {
		name []string
		data []string
	}
	tests := []struct {
		name     string
		fileMock fileMock
		want     map[string]string
	}{
		{
			name: "should parse single dictionary",
			fileMock: fileMock{
				name: []string{"foo.json"},
				data: []string{`{
					"vscode": "VS Code"
				}`},
			},
			want: map[string]string{
				"vscode": "VS Code",
			},
		},
		{
			name: "should combine dictionaries",
			fileMock: fileMock{
				name: []string{"foo.json", "bar.json"},
				data: []string{
					`{
						"vscode": "VS Code"
					}`,
					`{
						"wifi": "Wi-Fi"
					}`,
				},
			},
			want: map[string]string{
				"vscode": "VS Code",
				"wifi":   "Wi-Fi",
			},
		},
		{
			name: "should return empty dictionary on error",
			fileMock: fileMock{
				name: []string{"foo.json"},
				data: []string{"lorem ipsum"},
			},
			want: make(map[string]string),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fsys := fstest.MapFS{}

			for idx := range tc.fileMock.name {
				name := tc.fileMock.name[idx]
				data := tc.fileMock.data[idx]

				fsys[name] = &fstest.MapFile{
					Data: []byte(data),
				}
			}

			got := GetDictionary(fsys, tc.fileMock.name)

			assert.Equal(t, tc.want, got)
		})
	}
}
