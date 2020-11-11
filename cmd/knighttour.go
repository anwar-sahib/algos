package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	randomStart, warnsdorff bool
)

var chess [8][8]int
var moveX = [8]int{2, 1, -1, -2, -2, -1, 1, 2}
var moveY = [8]int{1, 2, 2, 1, -1, -2, -2, -1}
var N = 8

/*  Knight in middle of the chess board can have a maximum of 8 possible moves
         x axis
y  	_ |	6 |	_ |	7 | _
a	4 |	_ |	_ |	_ | 8
x	_ |	_ |	K | _ | _
i	4 |	_ |	_ |	_ |	1
s	_ | 3 | _ |	_ |	2


If k is at position (x,y) then each move will have follwing co-ordinates
1 (x+2,y+1),  2(x+1,y+2),  3(x-1,y+1),  4(x-2,y+1),  5(x-2,y-1),  6(x-1,y-2),  7(x+1,y-2),   8(x+2,y-1)
*/

func init() {
	knightTourCmd.PersistentFlags().BoolVarP(&randomStart, "random-start", "r", false, "Knight will have a random starting position")
	//In practice, Warnsdorf’s heuristic successfully finds a solution in linear time.
	knightTourCmd.PersistentFlags().BoolVarP(&warnsdorff, "use-warnsdorff", "w", false, "Use Warnsdorff’s algorithm that generates better heuristics")

	rootCmd.AddCommand(knightTourCmd)
}

var knightTourCmd = &cobra.Command{
	Use:   "knighttour",
	Short: "Provides a possible knight tour",
	Long:  "Prints the order in which knight travels in a chess board following its rules and visits every square",
	Run: func(cmd *cobra.Command, args []string) {
		startX := 0
		startY := 0
		if randomStart {
			startX = getGeneratedNumber(8)
			startY = getGeneratedNumber(8)
			fmt.Printf("Starting at position x=%d y=%d\n", startX, startY)
		}
		chess[startX][startY] = 1

		start := time.Now()

		fmt.Println("Calculating the tour, please wait...")
		kt(chess, startX, startY, 1)
		if showTime {
			elapsed := time.Since(start)
			fmt.Printf("Execution took %s\n", elapsed)
		}
	},
}

//There are N^2 Cells and for each, we have a maximum of 8 possible moves to choose from, so the worst running time is O(8^N^2).
func kt(chess [8][8]int, x int, y int, moveN int) bool {
	//All the squares are covered print the solution and return
	if moveN == 64 {
		fmt.Println()
		for i := 0; i < 8; i++ {
			for j := 0; j < 8; j++ {
				if chess[i][j] < 10 {
					fmt.Printf(" %d  ", chess[i][j])
				} else {
					fmt.Printf("%d  ", chess[i][j])
				}
			}
			fmt.Println()
		}
		return true
	} else {
		//Check for all possible moves
		for i := 0; i < 8; i++ {
			newX := x + moveX[i]
			newY := y + moveY[i]

			//Check if the selected move is valid
			if validMove(chess, newX, newY) {
				moveN++
				chess[newX][newY] = moveN

				//Recursively call the tour with new position
				if kt(chess, newX, newY, moveN) {
					return true
				}

				//Backtracking part
				moveN--
				chess[newX][newY] = 0
			}
		}
	}
	return false
}

func validMove(chess [8][8]int, x int, y int) bool {
	if (x >= 0 && x < N) &&
		(y >= 0 && y < N) &&
		chess[x][y] == 0 { //Square is unvisited
		return true
	}
	return false
}
