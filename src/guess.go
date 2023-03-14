package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var guessdigi = 30

	var count = 100

	for count > 0 {
		var digi = rand.Intn(100) + 1
		if guessdigi > digi {
			fmt.Printf("%v is smaller than my digi.\n", digi)
		} else if guessdigi < digi {
			fmt.Printf("%v is bigger than my digi.\n", digi)
		} else {
			fmt.Printf("You are right, it's %v.\n", digi)
			break
		}
		count--
	}
}
