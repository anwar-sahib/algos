package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

//A monad is an abstraction that allows structuring programs generically.

//The operation to be perfomed on the numbers
type operation func(int, int) int

//Perform function takes the operation to be performed, computes it and returns the result
type perform func(operation) int

/* pair is the generic defination that returns perform function on that pair of numbers.
The acutal operation to be done is decided in function that define this perform function.
Thus in the object-oriented paradigm you have an interface and two implementations, in functional you have a
function (interface) that accepts another function (implementation) and calls the implementation inside.
Thus here pair acts as interface. It is high order function that accepts two numbers and returns another function (perform),
which accepts the third function(implementation - add, sub, mul) that knows how to operate with this numbers.
*/
func pair(a, b int) perform {
	per := func(opr operation) int {
		return opr(a, b)
	}
	return per
}

var (
	sumOperation, subOperation, mulOperation bool
)

func init() {
	monadCmd.PersistentFlags().BoolVarP(&sumOperation, "add", "a", false, "Add two numbers")
	monadCmd.PersistentFlags().BoolVarP(&subOperation, "subtract", "s", false, "subtract two numbers")
	monadCmd.PersistentFlags().BoolVarP(&mulOperation, "multiply", "m", false, "Multiply two numbers")

	rootCmd.AddCommand(monadCmd)
}

var monadCmd = &cobra.Command{
	Use:   "monad",
	Short: "Demonstrates how to use function as variable in go",
	Long:  "Demonstrates how we can send a function as variable to other funciton in go lang",
	Run: func(cmd *cobra.Command, args []string) {
		a := 10
		b := 5
		fmt.Printf("a:%d  b:%d\n", a, b)

		if sumOperation {
			result := add(pair(a, b))
			fmt.Println("Addition:", result)
		}

		if subOperation {
			result := sub(pair(a, b))
			fmt.Println("Subtraction:", result)
		}
		if mulOperation {
			result := mul(pair(a, b))
			fmt.Println("Multiplication:", result)
		}
	},
}

//Implementations
func add(per perform) int {
	fnDef := func(a, b int) int {
		return a + b
	}
	return per(fnDef)
}

func sub(per perform) int {
	fnDef := func(a, b int) int {
		return a - b
	}
	return per(fnDef)
}

func mul(per perform) int {
	fnDef := func(a, b int) int {
		return a * b
	}
	return per(fnDef)
}
