package cmd

import (
	"fmt"
	"time"

	tm "github.com/buger/goterm"
	"github.com/spf13/cobra"
)

var chessSize int
var queenChess [100][100]int

func init() {
	nQueenProblemCmd.PersistentFlags().IntVarP(&chessSize, "size", "n", 1, "Number of cell in a row in the chess board")
	nQueenProblemCmd.PersistentFlags().BoolVarP(&showSteps, "show-steps", "s", false, "Shows intermediate steps for the magic square")

	rootCmd.AddCommand(nQueenProblemCmd)
}

var nQueenProblemCmd = &cobra.Command{
	Use:   "n-queen",
	Short: "Solves the n queen problem",
	Long:  "Solves the n queen problem for a chess board of size n * n",
	Run: func(cmd *cobra.Command, args []string) {
		ensureBelow(chessSize, 100, "cells")

		if showSteps {
			tm.Clear() // Clear current screen to show steps properly
		}

		res := nq(0, queenChess, showSteps)

		if !res {
			fmt.Printf("No solution for chess of size %d\n", chessSize)
		}
	},
}

//Each queen will be place in a new row. So q0 will be in row 0, q1 will be in row 1 and so on.
func nq(q int, queenChess [100][100]int, showSteps bool) bool {
	if q >= chessSize {
		printQueens(queenChess, chessSize, false)
		return true
	}

	for c := 0; c < chessSize; c++ {
		if isSafe(q, c, queenChess) {
			queenChess[q][c] = 1 //Set the queen at row q and column c

			if showSteps {
				printQueens(queenChess, chessSize, true)
			}

			if nq(q+1, queenChess, showSteps) { //Traverse to next queen
				return true
			}
			queenChess[q][c] = 0 //Backtrack the earlier step

			if showSteps {
				printQueens(queenChess, chessSize, true)
			}
		}
	}

	return false
}

func isSafe(row int, col int, chess [100][100]int) bool {
	//Check if any queen is in same column in some other row
	for r := 0; r < chessSize; r++ {
		if chess[r][col] == 1 { //This  check can be added as part of the below loop
			return false
		}
	}

	// Check if any queen is present on diagonals. As we are moving row wise,
	// we need to check diagonally only on rows before current position
	checkRow := 0
	checkCol := 0
	for i := 1; i < chessSize; i++ {
		checkRow = row - i
		checkCol = col - i

		if checkRow >= 0 && checkCol >= 0 {
			if chess[checkRow][checkCol] == 1 {
				return false
			}
		}

		checkRow = row - i
		checkCol = col + i

		if checkRow >= 0 && checkCol < chessSize {
			if chess[checkRow][checkCol] == 1 {
				return false
			}
		}
	}
	return true
}

func printQueens(queenChess [100][100]int, n int, flush bool) {
	tm.MoveCursor(1, 1)
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			fmt.Printf("%d  ", queenChess[x][y])
		}
		fmt.Println()
	}
	if flush {
		time.Sleep(time.Second)
		tm.Flush() // Call it every time at the end of rendering except last call
	}

}
