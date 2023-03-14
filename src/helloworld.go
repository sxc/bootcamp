package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	// Comments

	/*
		2 line
		comments
	*/
	fmt.Println("hello world")

	var a = "initial"
	const b string = "bike"
	var c float64 = 10000
	var d = 1
	const noon = 12
	const firstProgram = "Hello World!"
	g := a + "foo"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(math.Sin(c))
	fmt.Println(g)

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if num := -9; num < 0 {
		fmt.Println(num, "is negative")
	}

	now := time.Now()
	fmt.Println(now)

}
