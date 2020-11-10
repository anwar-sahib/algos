package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	expected := 12
	actual := add(pair(5, 7))

	assert.Equal(t, expected, actual)
}

func TestSub(t *testing.T) {
	expected := 4
	actual := sub(pair(12, 8))

	assert.Equal(t, expected, actual)
}

func TestMul(t *testing.T) {
	expected := 30
	actual := mul(pair(5, 6))

	assert.Equal(t, expected, actual)
}
