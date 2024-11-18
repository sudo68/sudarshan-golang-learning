package main

import "fmt"

// Constant initialization
const pi = 3.14

// Multiple constant initialization
const (
	big   = 1 << 100
	small = big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	const hello = "こんにちは、"
	fmt.Println(hello, "世界")
	fmt.Println("Happy", pi, "Day")
	fmt.Println(needInt(small))
	fmt.Println(needFloat(small))
	fmt.Println(needFloat(big))
}
