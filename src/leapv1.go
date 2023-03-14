package main

import "fmt"

func main() {
	const year = 2024
	if year%4 == 0 && year%100 != 0 {
		fmt.Printf(" %v is run", year)
	} else if year%400 == 0 {

		fmt.Printf(" %v is run", year)
	} else {
		fmt.Printf("%v Not Run", year)
	}
}
