package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var count = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		if rand.Intn(2) == 0 {
			break
		}
		count--
	}

	if count == 0 {
		fmt.Println("🚀Liftoff!")
	} else {
		fmt.Println("Launch failed. 💥")
	}

}
