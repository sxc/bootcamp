package main

import "fmt"

func main() {
	var weight = 149.0
	// weight = weight * 0.3783
	// weight *= 0.3783

	var age = 41
	age = age + 1
	age += 1
	age++
	weight = weight - 2
	fmt.Println(weight)
	fmt.Println(age)
}
