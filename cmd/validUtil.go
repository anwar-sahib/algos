package cmd

import (
	"fmt"
	"os"
)

const memocapacity = 100

func ensurePositive(number int, name string) {
	if number <= 0 {
		fmt.Printf("%s should be positive non zero integer\n", name)
		os.Exit(1)
	}
}

func ensureBelow(number int, max int, name string) {
	if number > max {
		fmt.Printf("%s value should not be greater than %d\n", name, max)
		os.Exit(1)
	}
}

func isOdd(number int) {
	if number%2 == 0 {
		fmt.Printf("%d should be odd\n", number)
		os.Exit(1)
	}
}
