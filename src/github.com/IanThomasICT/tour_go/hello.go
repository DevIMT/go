package main

import (
	"fmt"
)

func main() {
	var x int = 5
	var y int = 7
	var sum int = x + y

	fmt.Println(sum)

	var a[5] int
	a[2] = 7
	
	fmt.Println(a[2])

	// Short initialize syntax :=
	b := [5]int{5,4,3,2,1}			// Array
	c := []int{2,5,8,1,3,1}			// Slice of ints
	c = append(c,6)

	fmt.Println(b)
	fmt.Println(c)

	// Dictionary, map[key_type]value_type
	vertices := make(map[string]int)
	vertices["triangle"] = 3
	vertices["square"] = 4
	vertices["dodecagon"] = 12

	delete(vertices, "square")
	fmt.Println(vertices)

	// Go only uses for loops
	// Iterated for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While loop
	i := 10
	for i > 5 {
		fmt.Println(i)
		i--
	}

	arr := []string{"A","B","C"}
	mp := make(map[string]string)
	mp["a"] = "Alpha"
	mp["b"] = "Beta"
	// Range-based for loop (pythonic)
	for index, value := range arr {
		fmt.Println("index:",index,"value:",value)	// Println seperates comma-sep values by a space
	}

	for key, value := range mp {
		fmt.Println("index:",key,"value:",value)	// Println seperates comma-sep values by a space	
	}

	result := summation(9,12)
	fmt.Println("Sum: ",result)
}

// passed variables, return type at end
func summation(x int, y int) int {
	return x + y
}
