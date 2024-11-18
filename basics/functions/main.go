package main

import "fmt"

// Function to add numbers and return it.
func add(x int, y int) int {
	return x + y
}

// Function to add numbers but in shorthand syntax.
func add2(x, y int) int {
	return x + y
}

// Function returning multiple values.
func swap(x, y string) (string, string) {
	return y, x
}

// Naked returns in functions
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add2(43, 12))
	fmt.Println(swap("Mosh", "Moshi"))
	fmt.Println(split(17))
}
