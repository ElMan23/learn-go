package main

import (
	"fmt"
)

func main() {
	i := 7
	fmt.Println(i)
	fmt.Println(&i)

	inc(i)
	fmt.Println("With value", i)

	inc_ptr(&i)
	fmt.Println("With pointer", i)
}

func inc(x int) {
	x++
}

func inc_ptr(x *int) {
	* x++
}
