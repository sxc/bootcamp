package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const randrange = 401000000 - 56000000
	var num = rand.Intn(randrange) + 56000000 + 1
	fmt.Println(num)

}
