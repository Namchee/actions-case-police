package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeDictionary(t *testing.T) {
	type args struct {
		src  *map[string]string
		dest *map[string]string
	}

	tests := []struct {
		name string
		args args
		want *map[string]string
	}{
		{
			name: "should merge two map into one",
			args: args{
				src: &map[string]string{
					"foo": "bar",
				},
				dest: &map[string]string{
					"bar": "baz",
				},
			},
			want: &map[string]string{
				"foo": "bar",
				"bar": "baz",
			},
		},
		{
			name: "should replace map contents",
			args: args{
				src: &map[string]string{
					"foo": "bar",
				},
				dest: &map[string]string{
					"foo": "baz",
				},
			},
			want: &map[string]string{
				"foo": "baz",
			},
		},
		{
			name: "should not modify anything",
			args: args{
				src: &map[string]string{
					"foo": "bar",
				},
				dest: &map[string]string{},
			},
			want: &map[string]string{
				"foo": "bar",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			src := tc.args.src
			dest := tc.args.dest

			MergeDictionary(src, dest)

			assert.Equal(t, tc.want, src)
		})
	}
}

func TestRemoveEntries(t *testing.T) {
	type args struct {
		src       *map[string]string
		exclusion []string
	}

	tests := []struct {
		name string
		args args
		want *map[string]string
	}{
		{
			name: "should remove one of the key",
			args: args{
				src: &map[string]string{
					"github": "bar",
					"gitlab": "baz",
				},
				exclusion: []string{"github"},
			},
			want: &map[string]string{
				"gitlab": "baz",
			},
		},
		{
			name: "should do nothing",
			args: args{
				src: &map[string]string{
					"foo": "bar",
				},
				exclusion: []string{"bar"},
			},
			want: &map[string]string{
				"foo": "bar",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			src := tc.args.src
			exclude := tc.args.exclusion

			RemoveEntries(src, exclude)

			assert.Equal(t, tc.want, src)
		})
	}
}
