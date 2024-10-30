package main

import (
	"fmt"
	"math"
)

// A function to calculte square root of given number
func Sqrt(x float64) float64 {
	// Supplementary struct for calculations
	type SqrtHelper struct {
		now, prev float64
		treshold  float64
	}

	sqrtHelper := SqrtHelper{x, float64(0), 1e-10}

	// Square root calculation loop
	for math.Abs(sqrtHelper.now-sqrtHelper.prev) > sqrtHelper.treshold {
		sqrtHelper.prev = sqrtHelper.now
		sqrtHelper.now -= (sqrtHelper.now*sqrtHelper.now - x) / (2 * sqrtHelper.now)
	}
	return sqrtHelper.now
}

// Main function
func main() {
	fmt.Println(Sqrt(2))
}
