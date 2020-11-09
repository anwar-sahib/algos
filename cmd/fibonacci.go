package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	number                                                               int
	showSeries, dynamicApproach, dynamicRecursive, dynamicSpaceOptimized bool
	memo                                                                 []int
)
var fibonacciFn = fib

func init() {
	fibonacciCmd.PersistentFlags().IntVarP(&number, "number", "n", 1, "nth Number of fibonacci series")
	fibonacciCmd.PersistentFlags().BoolVarP(&showSeries, "show-series", "s", false, "Show the fibonacci series till nth number")
	fibonacciCmd.PersistentFlags().BoolVarP(&dynamicApproach, "dynamic", "d", false, "Use the dynamic approach (using array for memoization)")
	fibonacciCmd.PersistentFlags().BoolVarP(&dynamicRecursive, "dynamic-recursive", "r", false, "Use the dynamic approach along with recursive (using array for memoization)")
	fibonacciCmd.PersistentFlags().BoolVarP(&dynamicSpaceOptimized, "dynamic-space-optimized", "o", false, "Use the dynamic approach (using variables for memoization)")

	rootCmd.AddCommand(fibonacciCmd)
}

var fibonacciCmd = &cobra.Command{
	Use:   "fibonacci",
	Short: "Provides the nth fibonacci number",
	Long:  "Provides the value of nth number in the fibonacci series",
	Run: func(cmd *cobra.Command, args []string) {
		ensurePositive(number, "Number flag")
		if dynamicApproach {
			fibonacciFn = fibDp
			fmt.Println("Using dynamic approach (Non-recursive)")
		}

		if dynamicRecursive {
			fibonacciFn = fibDPRec
			memo = make([]int, number+1) // Initialize slice of memo, otherwise we get index out of range
			fmt.Println("Using dynamic approach (Recursive)")
		}

		if dynamicSpaceOptimized {
			fibonacciFn = fibDpSO
			fmt.Println("Using dynamic approach (Non-recursive) with space optimization")
		}

		start := time.Now()
		if showSeries {
			for i := 1; i <= number; i++ {
				fmt.Printf("%d ", fibonacciFn(i))
			}
			fmt.Println()
		} else {
			result := fibonacciFn(number)
			fmt.Println(result)
		}

		if showTime {
			elapsed := time.Since(start)
			fmt.Printf("Execution took %s\n", elapsed)
		}
	},
}

/* Simplest recusive soltion as this lends nicely for recursive programming. This is default implementation that will be called.
This has following issues,
1. For a very high number it leads to stack overflow error
2. The approach caculates same number fibnacci multiple time which it has already done. (This leads to long run time for large number)
Complexity : - O(2^n)
*/
func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

/*
Dynamic programming approach which saves the previously caculated results and re-uses it.
This makes the program highly efficient. Can be checked using the 'show-time' flag.
Time Complexity : - O(n)
Space Complexity : - O(n)
*/
func fibDp(n int) int {
	var f = make([]int, n+1)
	f[0] = 0
	f[1] = 1

	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

/*
Dynamic programming approach which saves the previously caculated results and re-uses it.
This makes the program more efficient but time is taken to access array location. Use the 'show-time' flag to check comparision.
AS this is recusive it still has problem of stack over flow in case of large number
Time Complexity : - O(n)
Space Complexity : - O(n)
*/
func fibDPRec(n int) int {
	if memo[n] != 0 {
		return memo[n]
	}
	if n == 1 || n == 2 {
		return 1
	} else {
		sum := fib(n-1) + fib(n-2)
		memo[n] = sum
		return sum
	}
}

/*
Dynamic programming approach which saves the previously caculated results and re-uses it using optimised space complexity.
As we need only the last two values, we are not saving everything in an array.
Time Complexity : - O(n)
Space Complexity : - O(1)
*/
func fibDpSO(n int) int {
	a := 1
	b := 1
	c := 1 //For case when n = 1
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}
