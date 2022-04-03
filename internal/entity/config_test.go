package entity

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfiguration(t *testing.T) {
	type expected struct {
		config *Configuration
		err    error
	}
	tests := []struct {
		name  string
		mocks map[string]string
		want  expected
	}{
		{
			name: "should read config correctly",
			mocks: map[string]string{
				"INPUT_GITHUB_TOKEN": "foo_bar",
				"INPUT_FIX":          "true",
				"INPUT_PRESET":       "abbreviates, brands",
				"INPUT_DICTIONARY":   "{\"foo\": \"bar\"}",
			},
			want: expected{
				config: &Configuration{
					Token:  "foo_bar",
					Fix:    true,
					Preset: []string{"abbreviates", "brands"},
					Dictionary: map[string]string{
						"foo": "bar",
					},
				},
				err: nil,
			},
		},
		{
			name:  "should throw an error when token is empty",
			mocks: map[string]string{},
			want: expected{
				config: nil,
				err:    errors.New("Missing GitHub token"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			for key, val := range tc.mocks {
				os.Setenv(key, val)
				defer os.Unsetenv(key)
			}

			got, err := ReadConfiguration()

			assert.Equal(t, tc.want.config, got)
			assert.Equal(t, tc.want.err, err)
		})
	}
}
