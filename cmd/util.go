package cmd

import (
	"math/rand"
	"time"
)

func getGeneratedArray(length int) []int {
	arr := make([]int, 0)
	//The environment in which these programs are executed is deterministic, so rand.Intn will always return the same number.
	source := rand.NewSource(time.Now().UnixNano()) //This helps generate random numbers
	random := rand.New(source)

	//Generate Random weights and values based on provided value
	valueRange := length * 5
	for i := 0; i < length; i++ {
		arr = append(arr, random.Intn(valueRange))
	}
	return arr
}

func getGeneratedNumber(maxValue int) int {
	//The environment in which these programs are executed is deterministic, so rand.Intn will always return the same number.
	source := rand.NewSource(time.Now().UnixNano()) //This helps generate random numbers
	random := rand.New(source)

	return random.Intn(maxValue)
}
