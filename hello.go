package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")

	var x int = 5
	var y int = 7

	var sum int = x + y

	fmt.Println(sum)

	z := 3
	fmt.Println(z)

	if z > 3 {
		fmt.Println("More than 3")
	} else if z < 3 {
		fmt.Println("Less than 3")
	} else {
		fmt.Println("Precisely 3")
	}

	var a  [5]int
	a[2] = 7
	fmt.Println(a)

	b := [3]int{1, 2, 3}
	fmt.Println(b)

	c := []int{3, 6, 9}
	fmt.Println(c)
	c = append(c, 12)
	fmt.Println(c)

	vertices := make(map[string]int)

	vertices["triangle"] = 3
	vertices["square"] = 4
	fmt.Println(vertices)
	fmt.Println(vertices["square"])
	delete(vertices, "square")
	fmt.Println(vertices)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}

	fmt.Println(sum(1, 2))
}

func sum(x int, y int) int {
	return x + y
}
