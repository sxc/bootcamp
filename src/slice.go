package main

import "fmt"

func main() {
	var x = [4]int{10, 20, 30}
	fmt.Println(x)
	x = [4]int{4, 5, 6, 7}
	fmt.Println(x)

	var y = []int{10, 20, 30}
	var z = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(y)
	fmt.Println(z)
	y = append(y, 40)
	fmt.Println(y)

	//var s = []int{}
	//data := []int{2, 4, 6, 8}

	fmt.Println(len(x))

}
