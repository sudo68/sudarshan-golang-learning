package main

import (
	"fmt"
	"math/cmplx"
)

var (
	// Boolean data types
	toBe bool = false
	// Unsigned integer data type
	maxInt uint64 = 1<<64 - 1
	// Complex data type
	z complex128 = cmplx.Sqrt(-5 + 2i)
)

func main() {
	fmt.Printf("Type: %T       Value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T     Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
