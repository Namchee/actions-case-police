package internal

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGithubClient(t *testing.T) {
	assert.NotPanics(t, func() {
		ctx := context.Background()

		NewGithubClient(ctx, "abcde")
	})
}
