package main

import (
	"fmt"
)

type exer4 int

func main() {
	var x exer4
	fmt.Print(x)
	fmt.Printf("%T", x)
	x = 42
	fmt.Print(x)
}
