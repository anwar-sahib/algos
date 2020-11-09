package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumInListEqualsN(t *testing.T) {
	numbers = []int{1, -2, 1, 0, 5}
	expected := true
	actual := sumInListEqualsN(6)

	assert.Equal(t, expected, actual)

	expected = false
	actual = sumInListEqualsN(8)
	assert.Equal(t, expected, actual)
}

func TestSumInListEqualsNOrderN(t *testing.T) {
	numbers = []int{11, 2, 7, -3, 4, 6, 1}
	expected := true
	actual := sumInListEqualsNOrderN(12)

	assert.Equal(t, expected, actual)

	expected = false
	actual = sumInListEqualsNOrderN(2)
	assert.Equal(t, expected, actual)
}
