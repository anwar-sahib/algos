package cmd

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

/*
Given a list of numbers and a number N, return whether any two numbers from the list add up to N.

For example, given [10, 15, -1, 3, -5, 7] and k of 17, return true since 10 + 7 is 17.
*/
var (
	sumValue, generateList int
	useHashSet             bool
)
var numbers = make([]int, 0)
var sumInListEqualsNFn = sumInListEqualsN

func init() {
	sumInListEqualsNCmd.PersistentFlags().IntVarP(&sumValue, "sum-value", "s", 0, "Value that should be equal to for sum of any two numbers in list")
	sumInListEqualsNCmd.PersistentFlags().IntVarP(&generateList, "generate-list", "g", 0, "Generate random list off numbers"+
		"(Should be less than "+strconv.Itoa(memocapacity)+")")
	sumInListEqualsNCmd.PersistentFlags().BoolVarP(&useHashSet, "performant", "p", false, "User this flag to find the answer in O(n)")

	rootCmd.AddCommand(sumInListEqualsNCmd)
}

var sumInListEqualsNCmd = &cobra.Command{
	Use:   "sumInListEqualsN",
	Short: "Checks if sum of any two numbers in a list equals a given value",
	Long:  "Checks if sum of any two numbers in a list including negative numbers equals given value",
	Run: func(cmd *cobra.Command, args []string) {

		if useHashSet {
			sumInListEqualsNFn = sumInListEqualsNOrderN
			fmt.Println("Using performant algorithm which provides answer in O(n)")

		}

		if generateList > 0 {
			//Check that the value is less than array capacity, as we are not using two dimensional array capacity is more
			ensureBelow(generateList, memocapacity*10, "generate-list")

			numbers = getGeneratedArray(generateList)
		} else {
			numbers = []int{1, -2, 1, 0, 5}
		}
		fmt.Println("Numbers list:", numbers)

		start := time.Now()
		sumInListEqualsNFn(sumValue)

		if showTime {
			elapsed := time.Since(start)
			fmt.Printf("Execution took %s\n", elapsed)
		}

	},
}

//Time Complexity - O(nlogn)
func sumInListEqualsN(sumValue int) bool {
	//This can be easily checked in O(n^2), but we will first sort the array to come up with better performance

	// It makes one call to data.Len to determine n, and O(n*log(n)) calls to data.Less and data.Swap.
	// The sort is not guaranteed to be stable.
	sort.Ints(numbers)

	left := 0
	right := len(numbers) - 1

	for left < right {
		if numbers[left]+numbers[right] == sumValue {
			fmt.Printf("%d and %d present in array add to %d\n", numbers[left], numbers[right], sumValue)
			return true
		} else if numbers[left]+numbers[right] < sumValue {
			left++
		} else {
			right--
		}
	}

	fmt.Printf("No numbers in array add to %d\n", sumValue)
	return false
}

/*HashSet provides a way to implement the search in O(n). However go does not provide Hashset and recommendation is
to use Map to implement Hashset. Hashset has best case O(1) and worst case O(n) when hash function maps everything to same bucket.
However we can use just keys to create a set with dummy values. This helps to have a set implementation of O(1) when key is integer.
A good explanation - https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics
*/
func sumInListEqualsNOrderN(sumValue int) bool {
	setImpl := make(map[int]bool)
	for i := 0; i < len(numbers); i++ {
		//Calculate the other numeber in pair that is required to get sum of sumValue
		otherNumber := sumValue - numbers[i]

		//As value is dummy boolean, we do not need it.
		_, present := setImpl[otherNumber]
		if present {
			fmt.Printf("%d and %d present in array add to %d\n", numbers[i], otherNumber, sumValue)
			return true //Skip this return in case we wish to print all pairs
		}

		//Setting a dummy true value, it does not matter what we set here
		setImpl[numbers[i]] = true
	}
	return false
}
