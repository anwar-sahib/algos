package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNq(t *testing.T) {
	chessSize = 6
	res := nq(0, queenChess)
	assert.True(t, res)

	chessSize = 3
	res = nq(0, queenChess)
	assert.False(t, res)

	chessSize = 12
	res = nq(0, queenChess)
	assert.True(t, res)
}
