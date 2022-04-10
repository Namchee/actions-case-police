package utils

import (
	"errors"
	"os"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

func TestReadEvent(t *testing.T) {
	type want struct {
		event int
		err   error
	}
	tests := []struct {
		name     string
		path     string
		mockFile []byte
		want     want
	}{
		{
			name:     "throw error when file cannot be read",
			path:     `/://///`,
			mockFile: []byte(`{ "foo": "bar" }`),
			want: want{
				event: 0,
				err:   errors.New("Failed to read event file"),
			},
		},
		{
			name:     "throw error when file cannot be parsed",
			path:     "/test.json",
			mockFile: []byte(`{ foo: "bar" }`),
			want: want{
				event: 0,
				err:   errors.New("Failed to parse event file"),
			},
		},
		{
			name:     "should return correctly on pull request events",
			path:     "/test.json",
			mockFile: []byte(`{ "action": "opened", "number": 1 }`),
			want: want{
				event: 1,
				err:   nil,
			},
		},
		{
			name:     "should return correctly on issue events",
			path:     "/test.json",
			mockFile: []byte(`{ "action": "opened", "issue": { "number": 2 } }`),
			want: want{
				event: 2,
				err:   nil,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv("GITHUB_EVENT_PATH", tc.path)
			defer os.Unsetenv("GITHUB_EVENT_PATH")

			mock := fstest.MapFS{
				tc.path[1:]: {
					Data: tc.mockFile,
				},
			}

			got, err := GetEventNumber(mock)

			assert.Equal(t, tc.want.event, got)
			assert.Equal(t, tc.want.err, err)
		})
	}
}
