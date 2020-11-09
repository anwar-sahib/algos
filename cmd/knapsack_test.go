package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKs(t *testing.T) {
	weight = []int{2, 3, 4, 3, 1}
	value = []int{5, 3, 11, 7, 2}
	expected := 25
	actual := ks(5, 10)

	assert.Equal(t, expected, actual)
}

func TestKsDp(t *testing.T) {
	weight = []int{3, 2, 5, 19}
	value = []int{19, 17, 6, 18}
	expected := 42
	actual := ksDp(4, 12)

	assert.Equal(t, expected, actual)
}

func TestKsDpPI(t *testing.T) {
	weight = []int{6, 23, 22, 16, 16, 8}
	value = []int{16, 26, 20, 13, 7, 13}
	expected := 29
	actual := ksPI(6, 15)

	assert.Equal(t, expected, actual)
}
