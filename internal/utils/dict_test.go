package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeDictionary(t *testing.T) {
	type args struct {
		a *map[string]string
		b *map[string]string
	}

	tests := []struct {
		name string
		args args
		want *map[string]string
	}{
		{
			name: "should merge two map into one",
			args: args{
				a: &map[string]string{
					"foo": "bar",
				},
				b: &map[string]string{
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
				a: &map[string]string{
					"foo": "bar",
				},
				b: &map[string]string{
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
				a: &map[string]string{
					"foo": "bar",
				},
				b: &map[string]string{},
			},
			want: &map[string]string{
				"foo": "bar",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a := tc.args.a
			b := tc.args.b

			MergeDictionary(a, b)

			assert.Equal(t, tc.want, a)
		})
	}
}
