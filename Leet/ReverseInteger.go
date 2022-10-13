package Leet

import (
	"fmt"
	"math"
)

func ReverseInteger(inputNumber int) int {

	fmt.Println("The input number is", inputNumber)
	signMultiplier := 1
	if inputNumber < 0 {
		signMultiplier = -1
		inputNumber = signMultiplier * inputNumber
	}

	var result int

	for inputNumber > 0 {
		// Add the current digit into result
		result = result*10 + inputNumber%10

		// Check if the result is within MaxInt32 and MinInt32 bounds
		if signMultiplier*result >= math.MaxInt32 || signMultiplier*result <= math.MinInt32 {
			return 0
		}
		inputNumber = inputNumber / 10

	}
	// Restore to original sign of number (+ or -)
	fmt.Println("The input number is", result)

	return signMultiplier * result
}
