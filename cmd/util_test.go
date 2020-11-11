package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGeneratedArray(t *testing.T) {
	actual := getGeneratedArray(6)
	assert.Equal(t, 6, len(actual))
}

func TestGetGeneratedNumber(t *testing.T) {
	number := getGeneratedNumber(10)
	assert.True(t, number <= 10)
}
