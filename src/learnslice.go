package main

import "fmt"

func main() {
	var x []int
	x = append(x, 10)
	fmt.Println(x)
	fmt.Println(len(x))
}
