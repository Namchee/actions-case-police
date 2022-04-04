package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMeta(t *testing.T) {
	type expected struct {
		meta *Meta
		err  error
	}
	tests := []struct {
		name string
		args string
		want expected
	}{
		{
			name: "should be able to extract metadata",
			args: "foo/bar",
			want: expected{
				meta: &Meta{
					Name:  "bar",
					Owner: "foo",
				},
				err: nil,
			},
		},
		{
			name: "should throw an error",
			args: "fake_github_repository",
			want: expected{
				meta: nil,
				err:  errors.New("Malformed GitHub metadata"),
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CreateMeta(tc.args)

			assert.Equal(t, tc.want.meta, got)
			assert.Equal(t, tc.want.err, err)
		})
	}
}
