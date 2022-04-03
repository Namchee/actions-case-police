package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsString(t *testing.T) {
	type args struct {
		input []string
		val   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "should return true",
			args: args{
				input: []string{"foo", "bar"},
				val:   "bar",
			},
			want: true,
		},
		{
			name: "should return false",
			args: args{
				input: []string{"foo", "bar"},
				val:   "baz",
			},
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := ContainsString(tc.args.input, tc.args.val)

			assert.Equal(t, tc.want, got)
		})
	}
}
