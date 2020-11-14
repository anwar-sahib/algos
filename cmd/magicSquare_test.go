package cmd

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMagicSquare(t *testing.T) {
	squareCells = 15
	createMagicSquare()
	isRowsCorrect(t)
	isColumnsCorrect(t)
}

func isRowsCorrect(t *testing.T) {
	fmt.Println("Sum of all elements in each row")
	for y := 0; y < squareCells; y++ {
		sumX := 0
		for x := 0; x < squareCells; x++ {
			sumX = sumX + square[x][y]
		}
		assert.Equal(t, 1695, sumX)
	}
}

func isColumnsCorrect(t *testing.T) {
	fmt.Println("Sum of all elements in each column")
	for x := 0; x < squareCells; x++ {
		sumY := 0
		for y := 0; y < squareCells; y++ {
			sumY = sumY + square[x][y]
		}
		assert.Equal(t, 1695, sumY)
	}
}
