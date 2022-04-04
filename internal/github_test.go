package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGithubClient(t *testing.T) {
	assert.NotPanics(t, func() {
		NewGithubClient("abcde")
	})
}
