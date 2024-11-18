package main

import "fmt"

// Variables declared with types.
var c, java, python bool
var i int

// With the presence of an intializer the variable with its type.
var i2, j2 = 1, 2
var c2, java2, python2 = true, false, "no!"

func main() {
	// Short hand variable declarations
	k, l := 3, 4
	c3, java3, python3 := true, true, "Yes!"

	fmt.Println(i, c, java, python)
	fmt.Printf("%d   %d   %#v %#v %s \n", i2, j2, c2, java2, python2)
	fmt.Printf("%T %T %T %T  %T \n", i2, j2, c2, java2, python2)
	fmt.Printf("%d %d\n", k, l)
	fmt.Printf("%v %v %s\n", c3, java3, python3)
}
