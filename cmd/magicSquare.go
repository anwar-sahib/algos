package cmd

import (
	"fmt"
	"time"

	tm "github.com/buger/goterm"
	"github.com/spf13/cobra"
)

var squareCells int
var square [100][100]int
var showSteps bool

func init() {
	magicSquareCmd.PersistentFlags().IntVarP(&squareCells, "no-rows", "n", 1, "Number of cell in a row in the magic square. (should be odd)")
	magicSquareCmd.PersistentFlags().BoolVarP(&showSteps, "show-steps", "s", false, "Shows intermediate steps for the magic square")

	rootCmd.AddCommand(magicSquareCmd)
}

var magicSquareCmd = &cobra.Command{
	Use:   "magic-square",
	Short: "Creates a magic square",
	Long:  "Creates a maigc sauare of nth cell",
	Run: func(cmd *cobra.Command, args []string) {
		isOdd(squareCells)
		ensureBelow(squareCells, 100, "cells")
		if showSteps {
			tm.Clear() // Clear current screen to show steps properly
		}
		createMagicSquare()

		if !showSteps {
			printSquare(false)
		}
	},
}

func createMagicSquare() {
	currentN := 1
	x := (squareCells - 1) / 2
	y := 0

	square[x][y] = currentN
	for currentN < squareCells*squareCells {
		nextX, nextY := next(x, y)
		currentN++
		if freeCell(nextX, nextY) {
			square[nextX][nextY] = currentN
			x, y = nextX, nextY

			if showSteps {
				printSquare(currentN != squareCells*squareCells)
			}
			continue
		}

		if exceedingX(nextX) && !exceedingY(nextY) {
			nextX = 0 //Traverse over x line
			square[nextX][nextY] = currentN
			x, y = nextX, nextY

			if showSteps {
				printSquare(currentN != squareCells*squareCells)
			}
			continue
		}

		if !exceedingX(nextX) && exceedingY(nextY) {
			nextY = squareCells - 1 //Traverse over y line
			square[nextX][nextY] = currentN
			x, y = nextX, nextY

			if showSteps {
				printSquare(currentN != squareCells*squareCells)
			}
			continue
		}

		if !freeCell(nextX, nextY) {
			nextX = nextX - 1 // Reversing the x step taken in next()
			nextY = nextY + 2 // Revesing the y step and sliding down

			square[nextX][nextY] = currentN
			x, y = nextX, nextY

			if showSteps {
				printSquare(currentN != squareCells*squareCells)
			}
			continue
		}

		if exceedingX(nextX) && exceedingY(nextY) {
			nextX = nextX - 1 // Reversing the x step taken in next()
			nextY = nextY + 2 // Revesing the y step and sliding down

			square[nextX][nextY] = currentN
			x, y = nextX, nextY

			if showSteps {
				printSquare(currentN != squareCells*squareCells)
			}
			continue
		}
	}
}

func next(x int, y int) (int, int) {
	return x + 1, y - 1
}

func freeCell(x int, y int) bool {
	if !exceedingX(x) && !exceedingY(y) {
		if square[x][y] != 0 {
			return false
		}
		return true
	}
	return false
}

func exceedingX(x int) bool {
	if x >= squareCells {
		return true
	}
	return false
}

func exceedingY(y int) bool {
	if y < 0 {
		return true
	}
	return false
}

func printSquare(flush bool) {
	tm.MoveCursor(1, 1)
	for y := 0; y < squareCells; y++ {
		for x := 0; x < squareCells; x++ {
			if square[x][y] < 10 {
				fmt.Printf("  %d   ", square[x][y])
			} else if square[x][y] < 100 {
				fmt.Printf(" %d   ", square[x][y])
			} else {
				fmt.Printf("%d   ", square[x][y])
			}
		}
		fmt.Println()
	}
	if flush {
		time.Sleep(time.Second)
		tm.Flush() // Call it every time at the end of rendering except last call
	}

}
