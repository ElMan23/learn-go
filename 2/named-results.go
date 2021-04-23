package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return	// This is a 'naked return', returns the named params x, y
}

func main() {
	fmt.Println(split(17))
}
