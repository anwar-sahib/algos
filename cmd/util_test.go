package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGeneratedArray(t *testing.T) {
	actual := getGeneratedArray(6)
	assert.Equal(t, 6, len(actual))
}
