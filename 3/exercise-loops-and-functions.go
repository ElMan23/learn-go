package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 2.0
	eps := 0.0001
	for z*z - x > eps {
		z = z - (z*z - x) / (2*x)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
