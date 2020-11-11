package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKt(t *testing.T) {
	var chess [8][8]int
	chess[0][0] = 1

	result := kt(chess, 0, 0, 1)
	assert.True(t, result)

}

func TestKtCustom(t *testing.T) {
	var chess [8][8]int
	chess[0][7] = 1

	result := kt(chess, 0, 7, 1)
	assert.True(t, result)

}
