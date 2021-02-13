package main

import (
	"fmt"
	"errors"
	"math"
)

func main() {

	var x float64 = 16
	result, err := sqrt(x)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Square root of", x, "is", result)
	}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined square for negative numbers")
	}
	return math.Sqrt(x), nil
}
