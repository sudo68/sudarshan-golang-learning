package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Printf("Value: %d   %d   %d\n", x, y, z)
	fmt.Printf("Type:  %T %T %T\n", x, y, z)
}
