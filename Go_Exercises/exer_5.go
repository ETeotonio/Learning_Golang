package main

import (
	"fmt"
)

type exer4 int

var x exer4
var y int

func main() {

	fmt.Print(x)
	fmt.Printf("%T", x)
	x = 42
	fmt.Print(x)

	fmt.Print(y)
	fmt.Printf("%T", y)
	y = int(x)
	fmt.Print(y)
}
