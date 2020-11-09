package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFib(t *testing.T) {
	expected := 8
	actual := fib(6)

	assert.Equal(t, expected, actual)
}

func TestFibDp(t *testing.T) {
	expected := 55
	actual := fibDp(10)

	assert.Equal(t, expected, actual)
}

func TestFibDpRec(t *testing.T) {
	memo = make([]int, 9)
	expected := 21
	actual := fibDPRec(8)

	assert.Equal(t, expected, actual)
}

func TestFibDpSO(t *testing.T) {
	expected := 75025
	actual := fibDpSO(25)

	assert.Equal(t, expected, actual)
}
