package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var (
	capacity, generateKs int
	showItems, dynamicks bool
)

var weight = make([]int, 0)
var value = make([]int, 0)
var memoks [memocapacity][memocapacity]int

var knapsackFn = ks

func init() {
	knapsackCmd.PersistentFlags().IntVarP(&capacity, "capacity", "c", 1, "Capacity of the Knapsack bag in terms of the weight")
	knapsackCmd.PersistentFlags().BoolVarP(&showItems, "show-items", "s", false, "Show the items present in the knapsack")
	knapsackCmd.PersistentFlags().BoolVarP(&dynamicks, "dynamic-approach", "d", false, "Use the dynamic approach (using array for memoization)")
	knapsackCmd.PersistentFlags().IntVarP(&generateKs, "generate-values", "g", 0, "Generate random weight and value pairs equal to the number"+
		"specified. (Should be less than "+strconv.Itoa(memocapacity)+")")

	rootCmd.AddCommand(knapsackCmd)
}

var knapsackCmd = &cobra.Command{
	Use:   "knapsack",
	Short: "Calculates the number of items that can be added in knapsack",
	Long:  "Provides the maximum number of items that can be added in a knapsack of given capacity such that the valus is maximum",
	Run: func(cmd *cobra.Command, args []string) {

		if dynamicks {
			knapsackFn = ksDp
			fmt.Println("Using dynamic approach (Recursive)")
		}

		if generateKs > 0 {
			//Check that the value is less than array capacity
			ensureBelow(generateKs, memocapacity, "generate-values")
			ensureBelow(capacity, memocapacity, "capacity")

			//Generate Random weights and values based on provided value
			weight = getGeneratedArray(generateKs)
			value = getGeneratedArray(generateKs)
		} else {
			weight = []int{2, 3, 4, 3, 1}
			value = []int{5, 3, 11, 7, 2}

		}
		fmt.Println("weight:", weight)
		fmt.Println("value: ", value)

		start := time.Now()
		result := knapsackFn(len(weight), capacity)
		fmt.Printf("Maximum value of %d can be accomodated in weight of %d\n", result, capacity)

		if showItems {
			knapsackFn = ksPI
			fmt.Println("Using bottom up approach (Non-recursive)")
			result := knapsackFn(len(weight), capacity)
			fmt.Printf("Maximum value of %d can be accomodated in weight of %d\n", result, capacity)
		}

		if showTime {
			elapsed := time.Since(start)
			fmt.Printf("Execution took %s\n", elapsed)
		}
	},
}

/* Dynamic approach with recusive programming
Time Complexity: O(2^n).
Auxiliary Space :O(1).
As no extra data structure has been used for storing values.
*/
func ks(n int, c int) int {
	result := 0
	if n == 0 || c == 0 {
		return 0
	} else if weight[n-1] > c {
		return ks(n-1, c)
	} else {
		option1 := ks(n-1, c)
		option2 := value[n-1] + ks(n-1, c-weight[n-1])
		result = maxInt(option1, option2)
	}
	return result
}

/* Dynamic approach with recusive programming using two dimensional array for memoization(as there can be possible n * c values)
Time Complexity: O(n*c).
Auxiliary Space :O(n*c).
As no extra data structure has been used for storing values.
*/
func ksDp(n int, c int) int {
	result := 0
	if n == 0 || c == 0 {
		return 0
	}

	if memoks[n-1][c] != 0 {
		return memoks[n-1][c]
	}

	if weight[n-1] > c {
		return ksDp(n-1, c)
	} else {
		option1 := ksDp(n-1, c)
		option2 := value[n-1] + ksDp(n-1, c-weight[n-1])
		result = maxInt(option1, option2)
		memoks[n-1][c] = result
	}
	return result
}

/* Non-recursive approach helps print the items selected. Need to find a way to do the same for recursive one.
Time Complexity: O(n*c).
Auxiliary Space :O(n*c).
*/
func ksPI(n int, c int) int {
	var k [memocapacity + 1][memocapacity + 1]int

	//Build the k in a bottom-up manner
	for item := 0; item <= n; item++ {
		for wt := 0; wt <= c; wt++ {
			//Finished with items or no more capacity
			if item == 0 || wt == 0 {
				k[item][wt] = 0
			} else if weight[item-1] <= wt {
				//Weight of next item is less than capacity, so we check both skipping of item and its selection
				k[item][wt] = maxInt(k[item-1][wt], value[item-1]+k[item-1][wt-weight[item-1]])
			} else {
				//Weight of next itme is more than capacity, so we skip and go to next item
				k[item][wt] = k[item-1][wt]
			}
			fmt.Printf("%d ", k[item][wt])
		}
		fmt.Println()
	}

	res := k[n][c]

	w := c
	for i := n; i > 0 && res > 0; i-- {

		// Backtrack from the result across each item, checking if it was selected or not
		if res == k[i-1][w] {
			//This implies the item was not selected
		} else {

			// This item is included.
			fmt.Printf("%d:%d \n", weight[i-1], value[i-1])

			// Since this weight is included, its value is deducted from result
			res = res - value[i-1]
			w = w - weight[i-1]
		}
	}
	return k[n][c]
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}
